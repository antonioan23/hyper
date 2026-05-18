package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"charm.land/catwalk/pkg/catwalk"
	"github.com/charmbracelet/hyper/internal/agent/cloudvio"
)

var cloudvioHTTPClient = &http.Client{Timeout: 30 * time.Second}

type cloudvioClient interface {
	GetModels(context.Context) ([]openaiModel, error)
}

type openaiModel struct {
	ID      string `json:"id"`
	OwnedBy string `json:"owned_by"`
}

var _ syncer[catwalk.Provider] = (*cloudvioSync)(nil)

type cloudvioSync struct {
	once       sync.Once
	result     catwalk.Provider
	cache      cache[catwalk.Provider]
	client     cloudvioClient
	baseURL    string
	autoupdate bool
	init       atomic.Bool
}

func (s *cloudvioSync) Init(client cloudvioClient, baseURL, path string, autoupdate bool) {
	s.client = client
	s.baseURL = baseURL
	s.cache = newCache[catwalk.Provider](path)
	s.autoupdate = autoupdate
	s.init.Store(true)
}

func (s *cloudvioSync) Get(ctx context.Context) (catwalk.Provider, error) {
	if !s.init.Load() {
		panic("called Get before Init")
	}

	var throwErr error
	s.once.Do(func() {
		if !s.autoupdate {
			slog.Info("Using embedded CloudVio provider")
			s.result = cloudvio.Embedded()
			return
		}

		cached, etag, cachedErr := s.cache.Get()
		if cached.ID == "" || cachedErr != nil {
			cached = cloudvio.Embedded()
		}

		slog.Info("Fetching CloudVio models")
		result, err := s.client.GetModels(ctx)
		if errors.Is(err, context.DeadlineExceeded) {
			slog.Warn("CloudVio models not updated in time")
			s.result = cached
			return
		}
		if err != nil {
			slog.Warn("Failed to fetch CloudVio models, using cached", "err", err)
			s.result = cached
			return
		}
		if len(result) == 0 {
			slog.Warn("CloudVio returned no models, using cached")
			s.result = cached
			return
		}

		provider := transformCloudvioModels(result, cached, s.baseURL, etag)
		s.result = provider
		throwErr = s.cache.Store(provider)
	})
	return s.result, throwErr
}

// transformCloudvioModels converts an OpenAI /v1/models response into a
// catwalk.Provider, preserving metadata from the cached provider when available.
func transformCloudvioModels(models []openaiModel, cached catwalk.Provider, baseURL, etag string) catwalk.Provider {
	embeddedMap := make(map[string]catwalk.Model, len(cached.Models))
	for _, m := range cached.Models {
		embeddedMap[m.ID] = m
	}

	result := make([]catwalk.Model, 0, len(models))
	for _, m := range models {
		if em, ok := embeddedMap[m.ID]; ok {
			result = append(result, em)
		} else {
			result = append(result, catwalk.Model{
				ID:               m.ID,
				Name:             formatModelName(m.ID),
				ContextWindow:    131072,
				DefaultMaxTokens: 16384,
			})
		}
	}

	apiEndpoint := baseURL + "/v1"
	if apiEndpoint == "/v1" {
		apiEndpoint = cloudvio.DefaultBaseURL + "/v1"
	}

	return catwalk.Provider{
		ID:                  cloudvio.Name,
		Name:                cloudvio.DisplayName,
		Type:                "openai-compat",
		APIEndpoint:         apiEndpoint,
		DefaultLargeModelID: cached.DefaultLargeModelID,
		DefaultSmallModelID: cached.DefaultSmallModelID,
		Models:              result,
		DefaultHeaders:      map[string]string{"X-CloudVio-ETag": etag},
	}
}

type realCloudvioClient struct {
	baseURL string
}

func (c realCloudvioClient) GetModels(ctx context.Context) ([]openaiModel, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/v1/models", nil)
	if err != nil {
		return nil, err
	}

	resp, err := cloudvioHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode == http.StatusNotModified {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var modelsResp struct {
		Data []openaiModel `json:"data"`
	}
	if err := json.Unmarshal(body, &modelsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal models response: %w", err)
	}

	return modelsResp.Data, nil
}

// UpdateCloudVio updates the CloudVio provider information from a specified source.
func UpdateCloudVio(pathOrURL string) error {
	var provider catwalk.Provider

	switch {
	case pathOrURL == "" || pathOrURL == "embedded":
		provider = cloudvio.Embedded()
		provider.APIEndpoint = cloudvio.DefaultBaseURL + "/v1"
	case strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://"):
		client := realCloudvioClient{baseURL: pathOrURL}
		models, err := client.GetModels(context.Background())
		if err != nil {
			return fmt.Errorf("failed to fetch models from CloudVio: %w", err)
		}
		embedded := cloudvio.Embedded()
		provider = transformCloudvioModels(models, embedded, pathOrURL, "")
	default:
		content, err := os.ReadFile(pathOrURL)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if err := json.Unmarshal(content, &provider); err != nil {
			return fmt.Errorf("failed to unmarshal provider data: %w", err)
		}
	}

	if err := newCache[catwalk.Provider](cachePathFor("cloudvio")).Store(provider); err != nil {
		return fmt.Errorf("failed to save CloudVio provider to cache: %w", err)
	}

	slog.Info("CloudVio provider updated successfully", "from", pathOrURL, "to", cachePathFor("cloudvio"))
	return nil
}

func formatModelName(id string) string {
	id = strings.TrimSuffix(id, "-free")
	parts := strings.Split(id, "-")
	for i, p := range parts {
		switch strings.ToLower(p) {
		case "gpt", "ai", "llm", "glm", "mimo", "omni", "tts", "voicedesign":
			parts[i] = strings.ToUpper(p)
		default:
			if len(p) > 0 {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
		}
	}
	return strings.Join(parts, " ")
}
