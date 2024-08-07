package create

import (
	"github.com/leonlarsson/bfstats-image-gen/canvas"
	"github.com/leonlarsson/bfstats-image-gen/shared"
	"github.com/leonlarsson/bfstats-image-gen/structs"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042Image(data structs.BF2042Data, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.RegularSkeletonType)

	canvas.DrawTimePlayed(ctx, data.Stats.TimePlayed)

	canvas.DrawStat1(ctx, data.Stats.Kills)
	canvas.DrawStat2(ctx, data.Stats.Deaths)

	canvas.DrawStat3(ctx, data.Stats.Assists)
	canvas.DrawStat4(ctx, data.Stats.Revives)

	canvas.DrawStat5BestClass(ctx, data.Stats.BestClass)

	canvas.DrawStat6(ctx, data.Stats.WlRatio)

	canvas.DrawRightStat1(ctx, data.Stats.KillsPerMatch)
	canvas.DrawRightStat2(ctx, data.Stats.KdRatio)
	canvas.DrawRightStat3(ctx, data.Stats.KillsPerMinute)
	canvas.DrawRightStat4Rank(ctx, data.Stats.Rank)

	return c, ctx
}
