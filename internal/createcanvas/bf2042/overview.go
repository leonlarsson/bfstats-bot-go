package bf2042

import (
	"github.com/leonlarsson/bfstats-go/internal/canvas"
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/shared"
	"github.com/leonlarsson/bfstats-go/internal/utils"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042OverviewImage(data shapes.GenericRegularData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.RegularSkeletonType)

	canvas.DrawStat1(ctx, data.Slots.L1)
	canvas.DrawStat2(ctx, data.Slots.L2)

	canvas.DrawStat3(ctx, data.Slots.L3)
	canvas.DrawStat4(ctx, data.Slots.L4)

	canvas.DrawBestClassImage(ctx, "BF2042", data.Slots.L5.Value)

	// If the best class is a BF2042 class, draw text to make room for the image. Otherwise, draw regular text.
	if utils.IsBaseBF2042Class("BF2042", data.Slots.L5.Value) {
		canvas.DrawStat5BestClass(ctx, data.Slots.L5)
	} else {
		canvas.DrawStat5(ctx, data.Slots.L5)
	}

	canvas.DrawStat6(ctx, data.Slots.L6)

	canvas.DrawRightStat1(ctx, data.Slots.R1)
	canvas.DrawRightStat2(ctx, data.Slots.R2)
	canvas.DrawRightStat3(ctx, data.Slots.R3)
	canvas.DrawRightStat4Rank(ctx, data.Slots.R4)

	return c, ctx
}
