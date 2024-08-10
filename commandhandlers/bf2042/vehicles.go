package commandhandlers

import (
	"errors"
	"sort"
	"strings"

	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/datafetchers/bf2042datafetcher"
	"github.com/leonlarsson/bfstats-bot-go/localization"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/utils"
	"github.com/tdewolff/canvas/renderers"
)

// HandleBF2042VehiclesCommand handles the bf2042 vehicles command.
func HandleBF2042VehiclesCommand(loc localization.LanguageLocalizer, platform, username string) error {
	data, err := bf2042datafetcher.FetchBF2042VehiclesData(platform, username)
	if err != nil {
		return err
	}

	if len(data.Data) < 9 {
		return errors.New(loc.Translate("messages/not_enough_vehicles", map[string]string{"vehicles": string(rune(len(data.Data)))}))
	}

	// Sort the vehicles by kills
	sort.Slice(data.Data, func(i, j int) bool {
		return data.Data[i].Stats.Kills.Value > data.Data[j].Stats.Kills.Value
	})

	// Build the vehicles slice
	var vehicles []canvasdatashapes.Stat
	for _, vehicle := range data.Data {
		vehicleSlice := canvasdatashapes.Stat{
			Name:  strings.TrimSpace(vehicle.Metadata.Name),
			Value: loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(vehicle.Stats.Kills.Value)}),
			Extra: loc.Translate("stats/title/x_kpm_and_time", map[string]string{"kpm": loc.FormatPercent(vehicle.Stats.KillsPerMinute.Value, 1), "time": vehicle.Stats.TimePlayed.DisplayValue}),
		}
		vehicles = append(vehicles, vehicleSlice)
	}

	// Create the image
	imageData := canvasdatashapes.BF2042VehiclesCanvasData{
		BaseData: canvasdatashapes.BaseData{
			Identifier: "BF2042-001",
			Username:   username,
			Platform:   utils.TRNPlatformNameToInt(platform),
			Meta: canvasdatashapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Vehicles",
			},
		},
		Vehicles: vehicles,
	}

	c, _ := create.CreateBF2042VehiclesImage(imageData, shared.SolidBackground)
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}

	return nil
}
