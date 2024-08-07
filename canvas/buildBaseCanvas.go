package canvas

import (
	"fmt"
	"strings"

	"github.com/leonlarsson/bfstats-image-gen/shared"
	"github.com/leonlarsson/bfstats-image-gen/structs"
	"github.com/leonlarsson/bfstats-image-gen/utils"
	"github.com/tdewolff/canvas"
)

func BuildBaseCanvas(game string, data structs.BaseData, skeletonType shared.SkeletonType) (*canvas.Canvas, *canvas.Context) {
	c, ctx := CreateStatsCanvasAndContext()

	game = strings.ToUpper(game)

	// Images
	DrawBackground(ctx, utils.GetRandomBackgroundImage(game, shared.SolidBackground), true)
	DrawSkeleton(ctx, skeletonType, shared.RegularStyle)
	DrawGameLogo(ctx, fmt.Sprintf("assets/images/%s/Logos/%s_LOGO_BG.png", game, game), shared.RegularStyle)

	// Identifier
	DrawIdentifier(ctx, data.Identifier)

	if skeletonType == shared.RegularSkeletonType {
		if game == "BF2042" /* TODO: AND if best class is base class (has an avatar) */ {
			DrawBestClassImage(ctx, "assets/images/BF2042/Specialists/Angel.png")
		} else if game != "BF2042" {
			// TODO: Draw slightly bigger best class image
		}

		DrawAvatar(ctx, "assets/images/DefaultGravatar.png")

		DrawUsernameRegular(ctx, data.Username)

	} else {

		DrawSegmentText(ctx, data.Meta.Segment)

		DrawUsernameGrid(ctx, data.Username)
	}

	DrawPlatformIcon(ctx, shared.Platform(data.Platform), skeletonType, false)

	DrawFooterWithText(ctx, "BY MOZZY", "BATTLEFIELDSTATS.COM", skeletonType)

	return c, ctx
}
