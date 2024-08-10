package create

import (
	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042VehiclesImage(data canvasdatashapes.BF2042VehiclesCanvasData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.GridSkeletonType)

	canvas.DrawAllGridStats(ctx, data.Vehicles)
	return c, ctx
}
