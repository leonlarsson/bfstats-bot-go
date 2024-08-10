package utils

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func IsBaseBF2042Class(game, class string) bool {
	game = strings.ToUpper(game)

	bf2042BaseClasses := []string{
		"Angel",
		"Boris",
		"Casper",
		"Dozer",
		"Falck",
		"Irish",
		"Mackay",
		"Paik",
		"Rao",
		"Sundance",
		"Lis",
		"Crawford",
		"Zain",
		"Blasco",
	}

	return game == "BF2042" && slices.Contains(bf2042BaseClasses, class)
}

func buildSLevelString(level int) string {
	level = level - 99
	if level < 10 {
		return fmt.Sprintf("S00%d", level)
	}
	if level < 100 {
		return fmt.Sprintf("S0%d", level)
	}
	if level < 1000 {
		return fmt.Sprintf("S%d", level)
	}

	return "S001"
}

func FormatRankString(rank int) string {
	return fmt.Sprintf("Rank %d (%s)", rank, buildSLevelString(rank))
}

func PercentileToString(percentile float64) string {
	if percentile == 0 {
		return ""
	}

	p := message.NewPrinter(language.English)
	adjustedPercentile := 100 - percentile

	// If the percentile is a whole number, don't show any decimal places
	if adjustedPercentile == math.Trunc(adjustedPercentile) {
		return "Top " + p.Sprintf("%.0f", adjustedPercentile) + "%"
	}

	// Otherwise, use one decimal place
	return "Top " + p.Sprintf("%.1f", adjustedPercentile) + "%"
}
