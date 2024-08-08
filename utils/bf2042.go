package utils

import (
	"slices"
	"strings"
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
