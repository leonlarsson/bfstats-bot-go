package bf2042

import (
	"errors"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-go/internal/canvas"
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/createcanvas/bf2042"
	"github.com/leonlarsson/bfstats-go/internal/datafetcher"
	"github.com/leonlarsson/bfstats-go/internal/datafetcher/types"
	"github.com/leonlarsson/bfstats-go/internal/localization"
	"github.com/leonlarsson/bfstats-go/internal/shared"
	"github.com/leonlarsson/bfstats-go/internal/utils"
)

// HandleBF2042VehiclesCommand handles the bf2042 vehicles command.
func HandleBF2042VehiclesCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate, loc localization.LanguageLocalizer) error {
	username, platform, usernameFailedValidation := utils.GetStatsCommandArgs(session, interaction, &loc)
	if usernameFailedValidation {
		return errors.New("username failed validation")
	}

	data, err := datafetcher.Fetch[types.TrnVehiclesResponse](utils.TRNBF2042VehiclesURL(platform, username))
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
	var vehicles []shapes.Slot
	for _, vehicle := range data.Data {
		vehicleSlice := shapes.Slot{
			Name:  strings.TrimSpace(vehicle.Metadata.Name),
			Value: loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(vehicle.Stats.Kills.Value)}),
			Extra: loc.Translate("stats/title/x_kpm_and_time", map[string]string{"kpm": loc.FormatFloat(vehicle.Stats.KillsPerMinute.Value), "time": vehicle.Stats.TimePlayed.DisplayValue}),
		}
		vehicles = append(vehicles, vehicleSlice)
	}

	// Create the image
	imageData := shapes.GenericGridData{
		BaseData: shapes.BaseData{
			Identifier: "BF2042-001",
			Username:   username,
			Platform:   utils.TRNPlatformNameToInt(platform),
			Meta: shapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Vehicles",
			},
		},
		Slots: vehicles,
	}

	c, _ := bf2042.CreateBF2042VehiclesImage(imageData, shared.SolidBackground)
	imgBuf := canvas.CanvasToBuffer(c)

	// Edit the response
	session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Files: []*discordgo.File{
			{
				Name:        "vehicles.png",
				ContentType: "image/png",
				Reader:      imgBuf,
			},
		},
	})

	return nil
}
