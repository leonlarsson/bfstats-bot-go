package bf2042

import (
	"errors"
	"sort"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/canvas"
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/createcanvas/bf2042"
	"github.com/leonlarsson/bfstats-go/internal/datafetcher"
	"github.com/leonlarsson/bfstats-go/internal/datafetcher/types"
	"github.com/leonlarsson/bfstats-go/internal/localization"
	"github.com/leonlarsson/bfstats-go/internal/shared"
	"github.com/leonlarsson/bfstats-go/internal/utils"
)

// HandleBF2042WeaponsCommand handles the bf2042 weapons command.
func HandleBF2042WeaponsCommand(interaction *events.ApplicationCommandInteractionCreate, loc localization.LanguageLocalizer) error {
	println("Bot: Handler running for command: bf2042weapons")
	username, platform, usernameFailedValidation := utils.GetStatsCommandArgs(interaction, loc)
	if usernameFailedValidation {
		return errors.New("username failed validation")
	}

	data, err := datafetcher.Fetch[types.TrnWeaponsResponse](utils.TRNBF2042WeaponsURL(platform, username), interaction, loc, username)
	if err != nil {
		return err
	}

	if len(data.Data) < 9 {
		return errors.New(loc.Translate("messages/not_enough_weapons", map[string]string{"weapons": string(rune(len(data.Data)))}))
	}

	// Sort the weapons by kills
	sort.Slice(data.Data, func(i, j int) bool {
		return data.Data[i].Stats.Kills.Value > data.Data[j].Stats.Kills.Value
	})

	// Build the weapons slice
	var weapons []shapes.Slot
	for _, weapon := range data.Data {
		weaponStat := shapes.Slot{
			Name:  strings.TrimSpace(weapon.Metadata.Name),
			Value: loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(weapon.Stats.Kills.Value)}),
			Extra: loc.Translate("stats/title/x_accuracy_and_kpm", map[string]string{"accuracy": loc.FormatPercent(weapon.Stats.ShotsAccuracy.Value), "kpm": loc.FormatFloat(weapon.Stats.KillsPerMinute.Value)}),
		}
		weapons = append(weapons, weaponStat)
	}

	// Create the image
	imageData := shapes.GenericGridData{
		BaseData: shapes.BaseData{
			Identifier: "BF2042-001",
			Username:   username,
			Platform:   utils.TRNPlatformNameToInt(platform),
			Meta: shapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Weapons",
			},
		},
		Slots: weapons,
	}

	c, _ := bf2042.CreateBF2042WeaponsImage(imageData, shared.SolidBackground)
	imgBuf := canvas.CanvasToBuffer(c)

	// Edit the response
	_, err = interaction.Client().Rest().UpdateInteractionResponse(interaction.ApplicationID(), interaction.Token(), discord.MessageUpdate{
		Files: []*discord.File{
			{
				Name:   "overview.png",
				Reader: imgBuf,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
