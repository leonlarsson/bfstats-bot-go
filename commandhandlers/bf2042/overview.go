package commandhandlers

import (
	"cmp"
	"fmt"
	"math"

	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	datafetcher "github.com/leonlarsson/bfstats-bot-go/datafetcher/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/structs"
	"github.com/leonlarsson/bfstats-bot-go/utils"
	"github.com/tdewolff/canvas/renderers"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func HandleBF2042OverviewCommand(platform, username string) error {
	data, err := datafetcher.FetchBF2042OverviewData(platform, username)
	if err != nil {
		return err
	}

	overviewSegment, err := utils.GetTRNSegmentByType(data.Data.Segments, "overview")
	if err != nil {
		return err
	}

	// Create the image
	imageData := structs.BF2042OverviewData{
		BaseData: structs.BaseData{
			Identifier: "BF2042-001",
			Username:   data.Data.PlatformInfo.PlatformUserHandle,
			Platform:   int(utils.TRNPlatformNameToInt(data.Data.PlatformInfo.PLatformSlug)),
			Avatar:     cmp.Or(data.Data.PlatformInfo.AvatarURL, "assets/images/DefaultGravatar.png"),
			Meta: structs.Meta{
				Game:    "Battlefield 2042",
				Segment: "Overview",
			},
		},
		Stats: structs.BF2042OverviewStats{
			TimePlayed: structs.Stat{
				Name:  "Time Played:",
				Value: overviewSegment.Stats.TimePlayed.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.TimePlayed.Percentile),
			},
			Kills: structs.Stat{
				Name:  "Kills:",
				Value: overviewSegment.Stats.Kills.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.Kills.Percentile),
			},
			Deaths: structs.Stat{
				Name:  "Deaths:",
				Value: overviewSegment.Stats.Deaths.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.Deaths.Percentile),
			},
			Assists: structs.Stat{
				Name:  "Assists:",
				Value: overviewSegment.Stats.Assists.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.Assists.Percentile),
			},
			Revives: structs.Stat{
				Name:  "Revives:",
				Value: overviewSegment.Stats.Revives.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.Revives.Percentile),
			},
			WlRatio: structs.Stat{
				Name:  "W/L Ratio:",
				Value: overviewSegment.Stats.WlPercentage.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.WlPercentage.Percentile),
			},
			// BestClass: structs.Stat{
			// 	Name:  "Best Class:",
			// 	Value: "Angel",
			// 	Extra: "2,813 kills | 15 hours",
			// },
			KillsPerMatch: structs.Stat{
				Name:  "Kills/Match:",
				Value: overviewSegment.Stats.KillsPerMatch.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.KillsPerMatch.Percentile),
			},
			KdRatio: structs.Stat{
				Name:  "K/D Ratio:",
				Value: overviewSegment.Stats.KdRatio.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.KdRatio.Percentile),
			},
			KillsPerMinute: structs.Stat{
				Name:  "Kills/Minute:",
				Value: overviewSegment.Stats.KillsPerMinute.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.KillsPerMinute.Percentile),
			},
			MultiKills: structs.Stat{
				Name:  "Multikills:",
				Value: overviewSegment.Stats.MultiKills.DisplayValue,
				Extra: percentileToString(overviewSegment.Stats.MultiKills.Percentile),
			},
			Rank: structs.RankStat{
				Name:    formatRankString(overviewSegment.Stats.Level.Value),
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

func percentileToString(percentile float64) string {
	if percentile == 0 {
		return ""
	}

	p := message.NewPrinter(language.English)
	adjustedPercentile := 100 - percentile

	// If the percentile is a whole number, don't show any decimal places
	if adjustedPercentile == math.Trunc(adjustedPercentile) {
		return "Top " + p.Sprintf("%.0f", adjustedPercentile) + "%"
	}

	// Otherwise, use one decimal place
	return "Top " + p.Sprintf("%.1f", adjustedPercentile) + "%"
}

func formatRankString(rank int) string {
	return fmt.Sprintf("Rank %d (%s)", rank, buildSLevelString(rank))
}

func buildSLevelString(level int) string {
	level = level - 99
	if level < 10 {
		return fmt.Sprintf("S00%d", level)
	}
	if level < 100 {
		return fmt.Sprintf("S0%d", level)
	}
	if level < 1000 {
		return fmt.Sprintf("S%d", level)
	}

	return "S001"
}
