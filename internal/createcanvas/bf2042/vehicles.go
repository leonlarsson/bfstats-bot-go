package bf2042

import (
	"github.com/leonlarsson/bfstats-go/internal/canvas"
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/shared"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042VehiclesImage(data shapes.GenericGridData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.GridSkeletonType)

	canvas.DrawAllGridSlots(ctx, data.Slots)
	return c, ctx
}
