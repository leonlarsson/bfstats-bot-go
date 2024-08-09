package commandhandlers

import (
	"fmt"

	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/datafetchers/bf2042datafetcher"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/utils"
	"github.com/tdewolff/canvas/renderers"
)

// HandleBF2042OverviewCommand handles the bf2042 overview command.
// TODO: In the future, this will take a loc param. Similar to bfstats-bot implementation.
func HandleBF2042OverviewCommand(platform, username string) error {
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
			Avatar:     utils.GetAvatarImageURL(data.Data.PlatformInfo.AvatarURL),
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
				Name:  "Kills:",
				Value: overviewSegment.Stats.Kills.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.Kills.Percentile),
			},
			Deaths: canvasdatashapes.Stat{
				Name:  "Deaths:",
				Value: overviewSegment.Stats.Deaths.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.Deaths.Percentile),
			},
			Assists: canvasdatashapes.Stat{
				Name:  "Assists:",
				Value: overviewSegment.Stats.Assists.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.Assists.Percentile),
			},
			Revives: canvasdatashapes.Stat{
				Name:  "Revives:",
				Value: overviewSegment.Stats.Revives.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.Revives.Percentile),
			},
			WlRatio: canvasdatashapes.Stat{
				Name:  "W/L Ratio:",
				Value: overviewSegment.Stats.WlPercentage.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.WlPercentage.Percentile),
			},
			// BestClass: canvasdatashapes.Stat{
			// 	Name:  "Best Class:",
			// 	Value: "Angel",
			// 	Extra: "2,813 kills | 15 hours",
			// },
			KillsPerMatch: canvasdatashapes.Stat{
				Name:  "Kills/Match:",
				Value: overviewSegment.Stats.KillsPerMatch.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMatch.Percentile),
			},
			KdRatio: canvasdatashapes.Stat{
				Name:  "K/D Ratio:",
				Value: overviewSegment.Stats.KdRatio.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.KdRatio.Percentile),
			},
			KillsPerMinute: canvasdatashapes.Stat{
				Name:  "Kills/Minute:",
				Value: overviewSegment.Stats.KillsPerMinute.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.KillsPerMinute.Percentile),
			},
			MultiKills: canvasdatashapes.Stat{
				Name:  "Multikills:",
				Value: overviewSegment.Stats.MultiKills.DisplayValue,
				Extra: utils.PercentileToString(overviewSegment.Stats.MultiKills.Percentile),
			},
			Rank: canvasdatashapes.RankStat{
				Name:    utils.FormatRankString(overviewSegment.Stats.Level.Value),
				Value:   fmt.Sprintf("%.0f%% to next rank", overviewSegment.Stats.LevelProgression.Value),
				RankInt: overviewSegment.Stats.Level.Value,
				Extra:   fmt.Sprintf("XP: %s", overviewSegment.Stats.XPAll.DisplayValue),
			},
		},
	}

	c, _ := create.CreateBF2042OverviewImage(imageData, shared.SolidBackground)
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}

	return nil
}
