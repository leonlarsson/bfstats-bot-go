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

	canvas.DrawL1(ctx, data.Slots.L1)
	canvas.DrawL2(ctx, data.Slots.L2)

	canvas.DrawL3(ctx, data.Slots.L3)
	canvas.DrawL4(ctx, data.Slots.L4)

	canvas.DrawBestClassImage(ctx, "BF2042", data.Slots.L5.Value)

	// If the best class is a BF2042 class, draw text to make room for the image. Otherwise, draw regular text.
	if utils.IsBaseBF2042Class("BF2042", data.Slots.L5.Value) {
		canvas.DrawL5BestClass(ctx, data.Slots.L5)
	} else {
		canvas.DrawL5(ctx, data.Slots.L5)
	}

	canvas.DrawL6(ctx, data.Slots.L6)

	canvas.DrawR1(ctx, data.Slots.R1)
	canvas.DrawR2(ctx, data.Slots.R2)
	canvas.DrawR3(ctx, data.Slots.R3)
	canvas.DrawR4Rank(ctx, data.Slots.R4)

	return c, ctx
}
