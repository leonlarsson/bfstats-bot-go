package create

import (
	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/structs"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042WeaponsImage(data structs.BF2042WeaponsData, style shared.BackgroundFormat) (*core.Canvas, *core.Context) {
	c, ctx := canvas.BuildBaseCanvas("BF2042", data.BaseData, shared.GridSkeletonType)

	canvas.DrawAllGridStats(ctx, data.Weapons)
	return c, ctx
}
