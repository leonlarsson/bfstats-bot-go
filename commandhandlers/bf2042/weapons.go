package commandhandlers

import (
	"errors"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/datafetchers/bf2042datafetcher"
	"github.com/leonlarsson/bfstats-bot-go/localization"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/utils"
)

// HandleBF2042WeaponsCommand handles the bf2042 weapons command.
func HandleBF2042WeaponsCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate, loc localization.LanguageLocalizer) error {
	username, platform, usernameFailedValidation := utils.GetStatsCommandArgs(session, interaction, &loc)
	if usernameFailedValidation {
		return errors.New("username failed validation")
	}

	data, err := bf2042datafetcher.FetchBF2042WeaponsData(platform, username)
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
	var weapons []canvasdatashapes.Stat
	for _, weapon := range data.Data {
		weaponStat := canvasdatashapes.Stat{
			Name:  strings.TrimSpace(weapon.Metadata.Name),
			Value: loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(weapon.Stats.Kills.Value)}),
			Extra: loc.Translate("stats/title/x_accuracy_and_kpm", map[string]string{"accuracy": loc.FormatPercent(weapon.Stats.ShotsAccuracy.Value), "kpm": loc.FormatFloat(weapon.Stats.KillsPerMinute.Value)}),
		}
		weapons = append(weapons, weaponStat)
	}

	// Create the image
	imageData := canvasdatashapes.BF2042WeaponsCanvasData{
		BaseData: canvasdatashapes.BaseData{
			Identifier: "BF2042-001",
			Username:   username,
			Platform:   utils.TRNPlatformNameToInt(platform),
			Meta: canvasdatashapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Weapons",
			},
		},
		Weapons: weapons,
	}

	c, _ := create.CreateBF2042WeaponsImage(imageData, shared.SolidBackground)
	imgBuf := canvas.CanvasToBuffer(c)

	// Edit the response
	session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Files: []*discordgo.File{
			{
				Name:        "weapons.png",
				ContentType: "image/png",
				Reader:      imgBuf,
			},
		},
	})

	return nil
}
