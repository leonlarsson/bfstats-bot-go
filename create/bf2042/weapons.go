package create

import (
	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042WeaponsImage(data canvasdatashapes.BF2042WeaponsCanvasData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.GridSkeletonType)

	canvas.DrawAllGridStats(ctx, data.Weapons)
	return c, ctx
}
