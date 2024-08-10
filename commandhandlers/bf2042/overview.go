package commandhandlers

import (
	"cmp"
	"fmt"

	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/datafetchers/bf2042datafetcher"
	"github.com/leonlarsson/bfstats-bot-go/localization"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/utils"
	"github.com/tdewolff/canvas/renderers"
)

// HandleBF2042OverviewCommand handles the bf2042 overview command.
// TODO: In the future, this will take a loc param. Similar to bfstats-bot implementation.
func HandleBF2042OverviewCommand(loc localization.LanguageLocalizer, platform, username string) error {
	data, err := bf2042datafetcher.FetchBF2042OverviewData(platform, username)
	if err != nil {
		return err
	}

	overviewSegment, err := utils.GetTRNSegmentByType(data.Data.Segments, "overview")
	if err != nil {
		return err
	}

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
			// BestClass: canvasdatashapes.Stat{
			// 	Name:  loc.TranslateWithColon("stats/title/bestclass"),
			// 	Value: "Angel",
			// 	Extra: "2,813 kills | 15 hours",
			// },
			KillsPerMatch: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kills_per_match"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMatch.Value, 1),
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMatch.Percentile),
			},
			KdRatio: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kd"),
				Value: loc.FormatFloat(overviewSegment.Stats.KdRatio.Value, 1),
				Extra: utils.PercentileToString(overviewSegment.Stats.KdRatio.Percentile),
			},
			KillsPerMinute: canvasdatashapes.Stat{
				Name:  loc.TranslateWithColon("stats/title/kpm"),
				Value: loc.FormatFloat(overviewSegment.Stats.KillsPerMinute.Value, 1),
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
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}

	return nil
}
