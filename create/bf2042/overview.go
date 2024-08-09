package create

import (
	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/utils"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042OverviewImage(data canvasdatashapes.BF2042OverviewCanvasData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.RegularSkeletonType)

	canvas.DrawTimePlayed(ctx, data.Stats.TimePlayed)

	canvas.DrawStat1(ctx, data.Stats.Kills)
	canvas.DrawStat2(ctx, data.Stats.Deaths)

	canvas.DrawStat3(ctx, data.Stats.Assists)
	canvas.DrawStat4(ctx, data.Stats.Revives)

	canvas.DrawBestClassImage(ctx, "BF2042", data.Stats.BestClass.Value)

	// 1. If the best class is a base class, draw the text with room for the image
	// 2. If the best class is "Unknown", draw the multi-kills stat
	// 3. If the best class is anything else ("BF3 Engineer, etc."), draw a regular stat
	if utils.IsBaseBF2042Class("BF2042", data.Stats.BestClass.Value) {
		canvas.DrawStat5BestClass(ctx, data.Stats.BestClass)
	} else if data.Stats.BestClass.Value == "Unknown" {
		canvas.DrawStat5(ctx, data.Stats.MultiKills)
	} else {
		canvas.DrawStat5(ctx, data.Stats.BestClass)
	}

	canvas.DrawStat6(ctx, data.Stats.WlRatio)

	canvas.DrawRightStat1(ctx, data.Stats.KillsPerMatch)
	canvas.DrawRightStat2(ctx, data.Stats.KdRatio)
	canvas.DrawRightStat3(ctx, data.Stats.KillsPerMinute)
	canvas.DrawRightStat4Rank(ctx, data.Stats.Rank)

	return c, ctx
}
