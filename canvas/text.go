package canvas

import (
	"github.com/leonlarsson/bfstats-image-gen/structs"
	"github.com/tdewolff/canvas"
)

var robotoFont = GetRobotoFontFamily()
var robotoMonoFont = GetRobotoMonoFontFamily()

var (
	statsTitleFace = robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontLight)
	statsValueFace = robotoFont.Face(PixelsToPoints(40), canvas.White, canvas.FontMedium)
	statsExtraFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)

	statsRankTitleFace = robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontMedium)
	statsRankValueFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)
	statsRankExtraFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)
)

// GetRobotoFontFamily returns the Roboto font family
func GetRobotoFontFamily() *canvas.FontFamily {
	font := canvas.NewFontFamily("Roboto")
	font.LoadFontFile("assets/fonts/Roboto-Medium.ttf", canvas.FontMedium)
	font.LoadFontFile("assets/fonts/Roboto-Light.ttf", canvas.FontLight)
	font.LoadFontFile("assets/fonts/Roboto-Thin.ttf", canvas.FontThin)
	return font
}

func GetRobotoMonoFontFamily() *canvas.FontFamily {
	font := canvas.NewFontFamily("Roboto Mono")
	font.LoadFontFile("assets/fonts/RobotoMono-Light.ttf", canvas.FontLight)
	return font
}

func GetFontsForLanguage(language string) (*canvas.FontFamily, *canvas.FontFamily) {
	if language == "en" {
		return robotoFont, robotoMonoFont
	}

	// More languages here

	return robotoFont, robotoMonoFont
}

func DrawIdentifier(ctx *canvas.Context, identifier string) {
	face := robotoMonoFont.Face(PixelsToPoints(13), canvas.RGBA(255, 255, 255, .2), canvas.FontLight)
	ctx.DrawText(1192, 20, canvas.NewTextLine(face, identifier, canvas.Right))
}

func DrawUsername(ctx *canvas.Context, username string) {
	face := robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontMedium)
	ctx.DrawText(950, 256, canvas.NewTextLine(face, username, canvas.Center))
}

func DrawTimePlayed(ctx *canvas.Context, timePlayed string) {
	face := robotoFont.Face(PixelsToPoints(30), canvas.White, canvas.FontLight)
	ctx.DrawText(950, 296, canvas.NewTextLine(face, timePlayed, canvas.Center))
}

func DrawStat1(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(57, 180, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(57, 221, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(57, 259, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat2(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(388, 180, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(388, 221, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(388, 259, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat3(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(57, 371, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(57, 412, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(57, 450, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat4(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(388, 371, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(388, 412, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(388, 450, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat5(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(57, 561, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(57, 602, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(57, 640, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat5BestClass(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(57, 561, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(116, 602, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(57, 640, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawStat6(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(388, 561, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(388, 602, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(388, 640, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Left))
}

func DrawRightStat1(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(723, 354, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(723, 395, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(1180, 354, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Right))
}

func DrawRightStat2(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(723, 451, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(723, 492, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(1180, 451, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Right))
}

func DrawRightStat3(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(723, 548, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(723, 589, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(1180, 548, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Right))
}

func DrawRightStat4(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(723, 658, canvas.NewTextLine(statsTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(723, 699, canvas.NewTextLine(statsValueFace, stat.Value, canvas.Left))
	ctx.DrawText(1180, 658, canvas.NewTextLine(statsExtraFace, stat.Extra, canvas.Right))
}

func DrawRightStat4Rank(ctx *canvas.Context, stat structs.Stat) {
	ctx.DrawText(723, 659, canvas.NewTextLine(statsRankTitleFace, stat.Name, canvas.Left))
	ctx.DrawText(723, 689, canvas.NewTextLine(statsRankValueFace, stat.Value, canvas.Left))
	ctx.DrawText(723, 721, canvas.NewTextLine(statsRankExtraFace, stat.Extra, canvas.Left))
}
