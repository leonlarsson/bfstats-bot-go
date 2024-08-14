package utils

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/leonlarsson/bfstats-go/internal/shared"
)

// GetRandomBackgroundImage returns a random background image from a given game and style
func GetRandomBackgroundImage(game string, style shared.BackgroundFormat) string {
	var filesContainingString []string

	game = strings.ToUpper(game)

	files, err := os.ReadDir(fmt.Sprintf("assets/images/%s/Backgrounds", game))
	if err != nil {
		return ""
	}

	searchString := "SOLID_BG"
	if style == shared.ImageBackground {
		searchString = "IMAGE_BG"
	}

	for _, file := range files {
		if strings.Contains(file.Name(), searchString) {
			filesContainingString = append(filesContainingString, file.Name())
		}
	}

	if len(files) == 0 {
		return ""
	}

	return fmt.Sprintf("assets/images/%s/Backgrounds/%s", game, filesContainingString[rand.IntN(len(filesContainingString))])
}
