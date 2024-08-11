package commandhandlers

import (
	"cmp"
	"errors"
	"fmt"
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

// HandleBF2042OverviewCommand handles the bf2042 overview command.
func HandleBF2042OverviewCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate, loc localization.LanguageLocalizer) error {
	username, platform, usernameFailedValidation := utils.GetStatsCommandArgs(session, interaction, &loc)
	if usernameFailedValidation {
		return errors.New("username failed validation")
	}

	data, err := bf2042datafetcher.FetchBF2042OverviewData(platform, username)
	if err != nil {
		return err
	}

	classData, err := bf2042datafetcher.FetchBF2042ClassesData(platform, username)
	if err != nil {
		return err
	}

	overviewSegment, err := utils.GetTRNSegmentByType(data.Data.Segments, "overview")
	if err != nil {
		return err
	}

	// Sort the classes by kills
	sort.Slice(classData.Data, func(i, j int) bool {
		return classData.Data[i].Stats.Kills.Value > classData.Data[j].Stats.Kills.Value
	})

	// Create the image
	imageData := canvasdatashapes.BF2042OverviewCanvasData{
		BaseData: canvasdatashapes.BaseData{
			Identifier: "BF2042-001",
			Username:   data.Data.PlatformInfo.PlatformUserHandle,
			Platform:   int(utils.TRNPlatformNameToInt(data.Data.PlatformInfo.PLatformSlug)),
			Avatar:     utils.CleanUserAvatar(cmp.Or(data.Data.PlatformInfo.AvatarURL, "assets/images/BF2042/Specialists/Angel.png")),
			Meta: canvasdatashapes.Meta{
				Game:    "Battlefield 2042",
				Segment: "Overview",
			},
		},
		Stats: canvasdatashapes.BF2042OverviewCanvasStats{
			TimePlayed: canvasdatashapes.Stat{
				Name:  "Time Played:",
				Value: overviewSegment.Stats.TimePlayed.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.TimePlayed.Percentile),
			},
			Kills: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kills"),
				Value: loc.FormatInt(overviewSegment.Stats.Kills.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Kills.Percentile),
			},
			Deaths: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/deaths"),
				Value: loc.FormatInt(overviewSegment.Stats.Deaths.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Deaths.Percentile),
			},
			Assists: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/assists"),
				Value: loc.FormatInt(overviewSegment.Stats.Assists.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Assists.Percentile),
			},
			Revives: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/revives"),
				Value: loc.FormatInt(overviewSegment.Stats.Revives.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.Revives.Percentile),
			},
			WlRatio: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/wlratio"),
				Value: overviewSegment.Stats.WlPercentage.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.WlPercentage.Percentile),
			},
			BestClass: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/bestclass"),
				Value: strings.TrimSpace(classData.Data[0].Metadata.Name),
				Extra: fmt.Sprintf("%s | %s", loc.Translate("stats/title/x_kills_short", map[string]string{"kills": loc.FormatInt(classData.Data[0].Stats.Kills.Value)}), classData.Data[0].Stats.TimePlayed.DisplayValue),
			},
			KillsPerMatch: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kills_per_match"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMatch.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMatch.Percentile),
			},
			KdRatio: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kd"),
				Value: fmt.Sprintf("%s (%s)", loc.FormatFloat(overviewSegment.Stats.KdRatio.Value), loc.FormatFloat(overviewSegment.Stats.HumanKdRatio.Value)),
				Extra: utils.PercentileToString(overviewSegment.Stats.KdRatio.Percentile),
			},
			KillsPerMinute: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kpm"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMinute.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMinute.Percentile),
			},
			MultiKills: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/multikills"),
				Value: loc.FormatInt(overviewSegment.Stats.MultiKills.Value),
				Extra: utils.PercentileToString(overviewSegment.Stats.MultiKills.Percentile),
			},
			Rank: canvasdatashapes.RankStat{
				Name:    utils.FormatRankString(overviewSegment.Stats.Level.Value),
				Value:   loc.Translate("stats/extra/percentage_to_next_rank", map[string]string{"percentage": fmt.Sprintf("%.0f%%", overviewSegment.Stats.LevelProgression.Value)}),
				RankInt: overviewSegment.Stats.Level.Value,
				Extra:   fmt.Sprintf("XP: %s", loc.FormatInt(overviewSegment.Stats.XPAll.Value)),
			},
		},
	}

	c, _ := create.CreateBF2042OverviewImage(imageData, shared.SolidBackground)
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
