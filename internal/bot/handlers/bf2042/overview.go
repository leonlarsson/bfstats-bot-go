package bf2042

import (
	"cmp"
	"errors"
	"fmt"
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

// HandleBF2042OverviewCommand handles the bf2042 overview command.
func HandleBF2042OverviewCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate, loc localization.LanguageLocalizer) error {
	username, platform, usernameFailedValidation := utils.GetStatsCommandArgs(session, interaction, &loc)
	if usernameFailedValidation {
		return errors.New("username failed validation")
	}

	overviewData, err := datafetcher.Fetch[types.TrnOverviewResponse](utils.TRNBF2042OverviewURL(platform, username))
	if err != nil {
		return err
	}

	classData, err := datafetcher.Fetch[types.TrnClassesResponse](utils.TRNBF2042ClassesURL(platform, username))
	if err != nil {
		return err
	}

	overviewSegment, err := utils.GetTRNSegmentByType(overviewData.Data.Segments, "overview")
	if err != nil {
		return err
	}

	// Sort the classes by kills
	sort.Slice(classData.Data, func(i, j int) bool {
		return classData.Data[i].Stats.Kills.Value > classData.Data[j].Stats.Kills.Value
	})

	// Create the image
	imageData := shapes.GenericRegularData{
		BaseData: shapes.BaseData{
			Identifier: "BF2042-001",
			Username:   overviewData.Data.PlatformInfo.PlatformUserHandle,
			Platform:   int(utils.TRNPlatformNameToInt(overviewData.Data.PlatformInfo.PLatformSlug)),
			Avatar:     utils.CleanUserAvatar(cmp.Or(overviewData.Data.PlatformInfo.AvatarURL, "assets/images/BF2042/Specialists/Angel.png")),
			Meta: shapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Overview",
			},
			TimePlayed: overviewSegment.Stats.TimePlayed.DisplayValue,
		},
		Slots: shapes.GenericRegularSlots{
			L1: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/kills"),
				Value: loc.FormatInt(overviewSegment.Stats.Kills.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Kills.Percentile),
			},
			L2: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/deaths"),
				Value: loc.FormatInt(overviewSegment.Stats.Deaths.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Deaths.Percentile),
			},
			L3: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/assists"),
				Value: loc.FormatInt(overviewSegment.Stats.Assists.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Assists.Percentile),
			},
			L4: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/revives"),
				Value: loc.FormatInt(overviewSegment.Stats.Revives.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Revives.Percentile),
			},
			L5: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/bestclass"),
				Value: strings.TrimSpace(classData.Data[0].Metadata.Name),
				Extra: fmt.Sprintf("%s | %s", loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(classData.Data[0].Stats.Kills.Value)}), classData.Data[0].Stats.TimePlayed.DisplayValue),
			},
			L6: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/wlratio"),
				Value: overviewSegment.Stats.WlPercentage.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.WlPercentage.Percentile),
			},
			R1: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/kills_per_match"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMatch.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMatch.Percentile),
			},
			R2: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/kd"),
				Value: fmt.Sprintf("%s (%s)", loc.FormatFloat(overviewSegment.Stats.KdRatio.Value), loc.FormatFloat(overviewSegment.Stats.HumanKdRatio.Value)),
				Extra: utils.PercentileToString(overviewSegment.Stats.KdRatio.Percentile),
			},
			R3: shapes.Slot{
				Name:  loc.TranslateWithColon("stats/title/kpm"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMinute.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMinute.Percentile),
			},
			R4: shapes.Slot{
				Name:  utils.FormatRankString(overviewSegment.Stats.Level.Value),
				Value: loc.Translate("stats/extra/percentage_to_next_rank", map[string]string{"percentage": fmt.Sprintf("%.0f%%", overviewSegment.Stats.LevelProgression.Value)}),
				Extra: fmt.Sprintf("XP: %s", loc.FormatInt(overviewSegment.Stats.XPAll.Value)),
			},
		},
	}

	// if the user is not a base BF2042 class, replace with another stat
	if !utils.IsBaseBF2042Class("BF2042", imageData.Slots.L5.Value) {
		imageData.Slots.L5 = shapes.Slot{
			Name:  loc.TranslateWithColon("stats/title/multikills"),
			Value: loc.FormatInt(overviewSegment.Stats.MultiKills.Value),
			Extra: utils.PercentileToString(overviewSegment.Stats.MultiKills.Percentile),
		}
	}

	c, _ := bf2042.CreateBF2042OverviewImage(imageData, shared.SolidBackground)
	imgBuf := canvas.CanvasToBuffer(c)

	// Edit the response
	session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Files: []*discordgo.File{
			{
				Name:        "overview.png",
				ContentType: "image/png",
				Reader:      imgBuf,
			},
		},
	})

	return nil
}
