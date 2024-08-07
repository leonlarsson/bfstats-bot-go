package canvas

import (
	"fmt"
	"strings"

	"github.com/leonlarsson/bfstats-image-gen/shared"
	"github.com/leonlarsson/bfstats-image-gen/utils"
	"github.com/tdewolff/canvas"
)

func BuildBaseCanvas(game string, useGridSkeleton shared.SkeletonType) (*canvas.Canvas, *canvas.Context) {
	c, ctx := CreateStatsCanvasAndContext()

	game = strings.ToUpper(game)

	font, _ := GetFontsForLanguage("en")

	// Images
	DrawBackground(ctx, utils.GetRandomBackgroundImage(game, shared.ImageBackground), true)
	DrawSkeleton(ctx, shared.SkeletonType(useGridSkeleton), shared.RegularStyle)
	DrawGameLogo(ctx, fmt.Sprintf("assets/images/%s/Logos/%s_LOGO_BG.png", game, game), shared.RegularStyle)

	// Identifier
	DrawIdentifier(ctx, "FECbLioP0ywuiztPUP")

	if useGridSkeleton == shared.RegularSkeletonType {
		if game == "BF2042" /* TODO: AND if best class is base class (has an avatar) */ {
			DrawBestClass(ctx, "assets/images/BF2042/Specialists/Angel.png")
		} else if game != "BF2042" {
			// TODO: Draw slightly bigger best class image
		}

		DrawAvatar(ctx, "assets/images/DefaultGravatar.png")
		DrawPlatformIcon(ctx, shared.PlatformPC)

		DrawUsername(ctx, "MozzyFX")
		DrawTimePlayed(ctx, "150 hours")

		DrawFooterWithText(ctx, "BY MOZZY", "BATTLEFIELDSTATS.COM")
	}

	return c, ctx
}
