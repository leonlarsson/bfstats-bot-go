package canvas

import (
	"fmt"
	"strings"

	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/shared"
	"github.com/leonlarsson/bfstats-go/internal/utils"
	"github.com/tdewolff/canvas"
)

func BuildBaseCanvas(game string, data shapes.BaseData, skeletonType shared.SkeletonType) (*canvas.Canvas, *canvas.Context) {
	c, ctx := CreateStatsCanvasAndContext()

	game = strings.ToUpper(game)

	// Images
	DrawBackground(ctx, utils.GetRandomBackgroundImage(game, shared.SolidBackground), true)
	DrawSkeleton(ctx, skeletonType, shared.RegularStyle)
	DrawGameLogo(ctx, fmt.Sprintf("assets/images/%s/Logos/%s_LOGO_BG.png", game, game), shared.RegularStyle)

	// Identifier
	DrawIdentifier(ctx, data.Identifier)

	if skeletonType == shared.RegularSkeletonType {
		DrawAvatar(ctx, data.Avatar)
		DrawUsernameRegular(ctx, data.Username)

	} else {
		DrawSegmentText(ctx, data.Meta.Segment)
		DrawUsernameGrid(ctx, data.Username)
	}

	DrawPlatformIcon(ctx, shared.Platform(data.Platform), skeletonType, false)

	DrawFooterWithText(ctx, "BY MOZZY", "BATTLEFIELDSTATS.COM", skeletonType)

	return c, ctx
}
