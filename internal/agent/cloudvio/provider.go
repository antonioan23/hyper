// Package cloudvio provides an embedded provider for the CloudVio
// OpenAI-compatible API proxy.
package cloudvio

import (
	_ "embed"
	"encoding/json"
	"log/slog"

	"charm.land/catwalk/pkg/catwalk"
)

//go:embed provider.json
var embedded []byte

// Embedded returns the embedded CloudVio provider.
var Embedded = func() catwalk.Provider {
	var provider catwalk.Provider
	if err := json.Unmarshal(embedded, &provider); err != nil {
		slog.Error("Could not use embedded CloudVio provider data", "err", err)
	}
	return provider
}

const (
	// Name is the identifier for the CloudVio provider.
	Name = "cloudvio"
	// DisplayName is the display name.
	DisplayName = "CloudVio"
	// DefaultBaseURL is the default CloudVio API endpoint.
	DefaultBaseURL = "http://179.42.8.135:8000/v1"
)
