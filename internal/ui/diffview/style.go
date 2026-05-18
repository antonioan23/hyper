package diffview

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// LineStyle defines the styles for a given line type in the diff view.
type LineStyle struct {
	LineNumber lipgloss.Style
	Symbol     lipgloss.Style
	Code       lipgloss.Style
}

// Style defines the overall style for the diff view, including styles for
// different line types such as divider, missing, equal, insert, and delete
// lines.
type Style struct {
	DividerLine LineStyle
	MissingLine LineStyle
	EqualLine   LineStyle
	InsertLine  LineStyle
	DeleteLine  LineStyle
	Filename    LineStyle
}

// DefaultLightStyle provides a default light theme style for the diff view.
func DefaultLightStyle() Style {
	return Style{
		DividerLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 77, G: 76, B: 87, A: 255}).
				Background(color.RGBA{R: 71, G: 118, B: 255, A: 255}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 135, G: 134, B: 127, A: 255}).
				Background(color.RGBA{R: 113, G: 154, B: 252, A: 255}),
		},
		MissingLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Background(color.RGBA{R: 223, G: 219, B: 221, A: 255}),
			Code: lipgloss.NewStyle().
				Background(color.RGBA{R: 223, G: 219, B: 221, A: 255}),
		},
		EqualLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 58, G: 57, B: 67, A: 255}).
				Background(color.RGBA{R: 223, G: 219, B: 221, A: 255}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 32, G: 31, B: 38, A: 255}).
				Background(color.RGBA{R: 241, G: 239, B: 239, A: 255}),
		},
		InsertLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 10, G: 220, B: 217, A: 255}).
				Background(lipgloss.Color("#c8e6c9")),
			Symbol: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 10, G: 220, B: 217, A: 255}).
				Background(lipgloss.Color("#e8f5e9")),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 32, G: 31, B: 38, A: 255}).
				Background(lipgloss.Color("#e8f5e9")),
		},
		DeleteLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 255, G: 56, B: 139, A: 255}).
				Background(lipgloss.Color("#ffcdd2")),
			Symbol: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 255, G: 56, B: 139, A: 255}).
				Background(lipgloss.Color("#ffebee")),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 32, G: 31, B: 38, A: 255}).
				Background(lipgloss.Color("#ffebee")),
		},
		Filename: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 77, G: 76, B: 87, A: 255}).
				Background(color.RGBA{R: 71, G: 118, B: 255, A: 255}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 77, G: 76, B: 87, A: 255}).
				Background(color.RGBA{R: 71, G: 118, B: 255, A: 255}),
		},
	}
}

// DefaultDarkStyle provides a default dark theme style for the diff view.
// Uses the Claude.com design system palette.
func DefaultDarkStyle() Style {
	return Style{
		DividerLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xff}).
				Background(color.RGBA{R: 0x38, G: 0x98, B: 0xec, A: 0xff}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xff}).
				Background(color.RGBA{R: 0x2a, G: 0x7a, B: 0xc8, A: 0xff}),
		},
		MissingLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Background(color.RGBA{R: 0x30, G: 0x30, B: 0x2e, A: 0xff}),
			Code: lipgloss.NewStyle().
				Background(color.RGBA{R: 0x30, G: 0x30, B: 0x2e, A: 0xff}),
		},
		EqualLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0x87, G: 0x86, B: 0x7f, A: 0xff}).
				Background(color.RGBA{R: 0x1e, G: 0x1e, B: 0x1c, A: 0xff}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xff}).
				Background(color.RGBA{R: 0x14, G: 0x14, B: 0x13, A: 0xff}),
		},
		InsertLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0x4e, G: 0x8c, B: 0x5e, A: 0xff}).
				Background(color.RGBA{R: 0x1f, G: 0x3d, B: 0x28, A: 0xff}),
			Symbol: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0x4e, G: 0x8c, B: 0x5e, A: 0xff}).
				Background(color.RGBA{R: 0x2d, G: 0x54, B: 0x37, A: 0xff}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xff}).
				Background(color.RGBA{R: 0x2d, G: 0x54, B: 0x37, A: 0xff}),
		},
		DeleteLine: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb5, G: 0x33, B: 0x33, A: 0xff}).
				Background(color.RGBA{R: 0x2d, G: 0x1a, B: 0x1a, A: 0xff}),
			Symbol: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb5, G: 0x33, B: 0x33, A: 0xff}).
				Background(color.RGBA{R: 0x3d, G: 0x22, B: 0x22, A: 0xff}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xff}).
				Background(color.RGBA{R: 0x3d, G: 0x22, B: 0x22, A: 0xff}),
		},
		Filename: LineStyle{
			LineNumber: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xff}).
				Background(color.RGBA{R: 0x38, G: 0x98, B: 0xec, A: 0xff}),
			Code: lipgloss.NewStyle().
				Foreground(color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xff}).
				Background(color.RGBA{R: 0x38, G: 0x98, B: 0xec, A: 0xff}),
		},
	}
}
