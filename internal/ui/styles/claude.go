package styles

import "image/color"

// Claude design system colors — dark terminal theme.
// Derived from the claude.com design system (Anthropic brand).
var (
	// Backgrounds (darkest to lightest).
	claudeNearBlack = color.RGBA{R: 0x14, G: 0x14, B: 0x13, A: 0xFF} // #141413 — primary dark bg
	claudeDarkWarm  = color.RGBA{R: 0x1e, G: 0x1e, B: 0x1c, A: 0xFF} // #1e1e1c — bgLeastVisible
	claudeDarkSrf   = color.RGBA{R: 0x30, G: 0x30, B: 0x2e, A: 0xFF} // #30302e — bgLessVisible, separator
	claudeIronWarm  = color.RGBA{R: 0x4d, G: 0x4c, B: 0x48, A: 0xFF} // #4d4c48 — bgMostVisible

	// Foregrounds (brightest to dimmest).
	claudeParchment  = color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xFF} // #f5f4ed — fgBase, onPrimary
	claudeWarmSilver = color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xFF} // #b0aea5 — fgSubtle
	claudeStoneGray  = color.RGBA{R: 0x87, G: 0x86, B: 0x7f, A: 0xFF} // #87867f — fgMoreSubtle
	claudeOliveGray  = color.RGBA{R: 0x5e, G: 0x5d, B: 0x59, A: 0xFF} // #5e5d59 — fgMostSubtle

	// Brand.
	claudeTerracotta = color.RGBA{R: 0xc9, G: 0x64, B: 0x42, A: 0xFF} // #c96442 — primary brand
	claudeCoral      = color.RGBA{R: 0xd9, G: 0x77, B: 0x57, A: 0xFF} // #d97757 — secondary, keyword
	claudeCoralLight = color.RGBA{R: 0xe8, G: 0xa6, B: 0x7a, A: 0xFF} // #e8a67a — accent
	claudeApricot    = color.RGBA{R: 0xf0, G: 0xbb, B: 0x8e, A: 0xFF} // #f0bb8e — warm highlight

	// Status.
	claudeErrorCrimson  = color.RGBA{R: 0xb5, G: 0x33, B: 0x33, A: 0xFF} // #b53333 — destructive, error
	claudeWarningGold   = color.RGBA{R: 0xc9, G: 0xa2, B: 0x4d, A: 0xFF} // #c9a24d — warning
	claudeWarningBright = color.RGBA{R: 0xda, G: 0xb5, B: 0x55, A: 0xFF} // #dab555 — warning subtle
	claudeBusyCoral     = color.RGBA{R: 0xd9, G: 0x77, B: 0x57, A: 0xFF} // #d97757 — busy (same as coral)
	claudeFocusBlue     = color.RGBA{R: 0x38, G: 0x98, B: 0xec, A: 0xFF} // #3898ec — info
	claudeBlueMid       = color.RGBA{R: 0x2a, G: 0x7a, B: 0xc8, A: 0xFF} // #2a7ac8 — info more subtle
	claudeBlueDeep      = color.RGBA{R: 0x1f, G: 0x5c, B: 0x99, A: 0xFF} // #1f5c99 — info most subtle
	claudeForest        = color.RGBA{R: 0x4e, G: 0x8c, B: 0x5e, A: 0xFF} // #4e8c5e — success
	claudeForestMid     = color.RGBA{R: 0x3d, G: 0x70, B: 0x49, A: 0xFF} // #3d7049 — success more subtle
	claudeForestDeep    = color.RGBA{R: 0x2d, G: 0x54, B: 0x37, A: 0xFF} // #2d5437 — success most subtle

	// Syntax highlighting.
	claudeSyntaxKeyword   = color.RGBA{R: 0xd9, G: 0x77, B: 0x57, A: 0xFF} // #d97757 — coral for keywords
	claudeSyntaxType      = color.RGBA{R: 0xe8, G: 0xa6, B: 0x7a, A: 0xFF} // #e8a67a — types
	claudeSyntaxString    = color.RGBA{R: 0xc9, G: 0xa2, B: 0x4d, A: 0xFF} // #c9a24d warm gold for strings
	claudeSyntaxFunc      = color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xFF} // parchment for functions
	claudeSyntaxComment   = color.RGBA{R: 0x5e, G: 0x5d, B: 0x59, A: 0xFF} // olive gray for comments
	claudeSyntaxBuiltin   = color.RGBA{R: 0xe8, G: 0xa6, B: 0x7a, A: 0xFF} // builtins
	claudeSyntaxTag       = color.RGBA{R: 0xc9, G: 0x64, B: 0x42, A: 0xFF} // terracotta for tags
	claudeSyntaxAttr      = color.RGBA{R: 0xd9, G: 0x77, B: 0x57, A: 0xFF} // attributes
	claudeSyntaxOperator  = color.RGBA{R: 0xb0, G: 0xae, B: 0xa5, A: 0xFF} // operators
	claudeSyntaxClass     = color.RGBA{R: 0xf5, G: 0xf4, B: 0xed, A: 0xFF} // class names
	claudeSyntaxDecorator = color.RGBA{R: 0xc9, G: 0xa2, B: 0x4d, A: 0xFF} // decorators
	claudeSyntaxNumber    = color.RGBA{R: 0xc9, G: 0x64, B: 0x42, A: 0xFF} // terracotta for numbers

	// Diff backgrounds (subtle tints).
	claudeDiffAddBg     = color.RGBA{R: 0x2d, G: 0x54, B: 0x37, A: 0xFF} // #2d5437 — addition line bg
	claudeDiffDelBg     = color.RGBA{R: 0x3d, G: 0x22, B: 0x22, A: 0xFF} // #3d2222 — deletion line bg
	claudeDiffOldHdr    = color.RGBA{R: 0x4d, G: 0x2a, B: 0x2a, A: 0xFF} // #4d2a2a — old header bg
	claudeDiffNewHdr    = color.RGBA{R: 0x2d, G: 0x54, B: 0x37, A: 0xFF} // #2d5437 — new header bg
	claudeDiffAddSubtle = color.RGBA{R: 0x1f, G: 0x3d, B: 0x28, A: 0xFF} // #1f3d28 — subtle addition bg
	claudeDiffDelSubtle = color.RGBA{R: 0x2d, G: 0x1a, B: 0x1a, A: 0xFF} // #2d1a1a — subtle deletion bg
)
