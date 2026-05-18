package styles

// ThemeForProvider returns the Styles associated with the given provider ID.
func ThemeForProvider(providerID string) Styles {
	return HyperTheme()
}

// HyperTheme returns the default Claude-themed Styles.
func HyperTheme() Styles {
	return quickStyle(quickStyleOpts{
		primary:   claudeTerracotta,
		secondary: claudeCoral,
		accent:    claudeCoralLight,
		keyword:   claudeCoral,

		fgBase:       claudeParchment,
		fgMoreSubtle: claudeStoneGray,
		fgSubtle:     claudeWarmSilver,
		fgMostSubtle: claudeOliveGray,

		onPrimary: claudeParchment,

		bgBase:         claudeNearBlack,
		bgLeastVisible: claudeDarkWarm,
		bgLessVisible:  claudeDarkSrf,
		bgMostVisible:  claudeIronWarm,

		separator: claudeDarkSrf,

		destructive:       claudeErrorCrimson,
		error:             claudeErrorCrimson,
		warningSubtle:     claudeWarningBright,
		warning:           claudeWarningGold,
		busy:              claudeBusyCoral,
		info:              claudeFocusBlue,
		infoMoreSubtle:    claudeBlueMid,
		infoMostSubtle:    claudeBlueDeep,
		success:           claudeForest,
		successMoreSubtle: claudeForestMid,
		successMostSubtle: claudeForestDeep,
	})
}

// Deprecated: Use HyperTheme() instead.
func CharmtonePantera() Styles {
	return HyperTheme()
}

// Deprecated: Use HyperTheme() instead.
func HyperhyperObsidiana() Styles {
	return HyperTheme()
}
