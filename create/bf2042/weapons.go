package create

import (
	"github.com/leonlarsson/bfstats-bot-go/internal/canvas"
	shapes "github.com/leonlarsson/bfstats-bot-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-bot-go/internal/shared"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042WeaponsImage(data shapes.BF2042WeaponsCanvasData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.GridSkeletonType)

	canvas.DrawAllGridStats(ctx, data.Weapons)
	return c, ctx
}
