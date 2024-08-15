package canvas

import (
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/tdewolff/canvas"
)

var robotoFont = GetRobotoFontFamily()
var robotoMonoFont = GetRobotoMonoFontFamily()

var (
	slotTitleFace = robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontLight)
	slotValueFace = robotoFont.Face(PixelsToPoints(40), canvas.White, canvas.FontMedium)
	slotExtraFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)

	slotRankTitleFace = robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontMedium)
	slotRankValueFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)
	slotRankExtraFace = robotoFont.Face(PixelsToPoints(25), canvas.White, canvas.FontLight)
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

func DrawSegmentText(ctx *canvas.Context, segment string) {
	face := robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontLight)
	textLine := canvas.NewTextLine(face, segment, canvas.Center)

	bounds := textLine.Bounds()

	xPadding := 60.0
	yPadding := 20.0
	bounds.W += xPadding
	bounds.H += yPadding

	ctx.SetFillColor(canvas.RGBA(32, 32, 32, 0.65))

	// Draw segment background
	ctx.DrawPath(600-(xPadding/2), 42-(yPadding/4), bounds.ToPath())
	ctx.DrawText(600, 69, textLine)
}

func DrawUsernameRegular(ctx *canvas.Context, username string) {
	face := robotoFont.Face(PixelsToPoints(35), canvas.White, canvas.FontMedium)
	ctx.DrawText(950, 256, canvas.NewTextLine(face, username, canvas.Center))
}

func DrawUsernameGrid(ctx *canvas.Context, username string) {
	face := robotoFont.Face(PixelsToPoints(30), canvas.White, canvas.FontMedium)
	textLine := canvas.NewTextLine(face, username, canvas.Right)
	textLineWidth := textLine.Width

	bounds := textLine.Bounds()

	xPadding := 20.0

	// Add padding
	bounds.W += textLineWidth + xPadding
	bounds.H += 10

	ctx.SetFillColor(canvas.RGBA(32, 32, 32, 0.65))

	// Draw username background
	// X is text x - text width - half the padding
	ctx.DrawPath((1100 - textLineWidth - (xPadding / 2)), 42, bounds.ToPath())
	ctx.DrawText(1100, 67, textLine)
}

func DrawTimePlayed(ctx *canvas.Context, timePlayed string) {
	face := robotoFont.Face(PixelsToPoints(30), canvas.White, canvas.FontLight)
	ctx.DrawText(950, 296, canvas.NewTextLine(face, timePlayed, canvas.Center))
}

func DrawL1(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(57, 180, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(57, 221, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 259, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL2(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(388, 180, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(388, 221, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(388, 259, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL3(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(57, 371, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(57, 412, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 450, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL4(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(388, 371, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(388, 412, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(388, 450, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL5(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(57, 561, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(57, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL5BestClass(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(57, 561, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(116, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawL6(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(388, 561, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(388, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(388, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawR1(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(723, 354, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(723, 395, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(1180, 354, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Right))
}

func DrawR2(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(723, 451, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(723, 492, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(1180, 451, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Right))
}

func DrawR3(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(723, 548, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(723, 589, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(1180, 548, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Right))
}

func DrawR4(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(723, 658, canvas.NewTextLine(slotTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(723, 699, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(1180, 658, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Right))
}

func DrawR4Rank(ctx *canvas.Context, slot shapes.Slot) {
	ctx.DrawText(723, 659, canvas.NewTextLine(slotRankTitleFace, slot.Name, canvas.Left))
	ctx.DrawText(723, 689, canvas.NewTextLine(slotRankValueFace, slot.Value, canvas.Left))
	ctx.DrawText(723, 721, canvas.NewTextLine(slotRankExtraFace, slot.Extra, canvas.Left))
}

/* Grid Slots */

// GridSlotTextBox returns a new text box with the given text. This is used to avoid overflowing text.
func GridSlotTextBox(text string) *canvas.Text {
	// 335 is the width of the text box
	// 50 is the height of the text box (to only allow one line, no wrapping)
	return canvas.NewTextBox(slotTitleFace, text, 335, 50, 0, 0, 0, 0)
}

func DrawGridSlot1(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(57, 180+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(57, 221, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 259, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot2(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(439, 180+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(439, 221, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(439, 259, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot3(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(823, 180+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(823, 221, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(823, 259, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot4(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(57, 371+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(57, 412, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 450, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot5(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(439, 371+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(439, 412, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(439, 450, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot6(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(823, 371+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(823, 412, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(823, 450, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot7(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(57, 561+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(57, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(57, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot8(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(439, 561+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(439, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(439, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawGridSlot9(ctx *canvas.Context, slot shapes.Slot) {
	nameTextBox := GridSlotTextBox(slot.Name)
	ctx.DrawText(823, 561+nameTextBox.Bounds().Y+9, nameTextBox)
	ctx.DrawText(823, 602, canvas.NewTextLine(slotValueFace, slot.Value, canvas.Left))
	ctx.DrawText(823, 640, canvas.NewTextLine(slotExtraFace, slot.Extra, canvas.Left))
}

func DrawAllGridSlots(ctx *canvas.Context, slots []shapes.Slot) {
	DrawGridSlot1(ctx, slots[0])
	DrawGridSlot2(ctx, slots[1])
	DrawGridSlot3(ctx, slots[2])
	DrawGridSlot4(ctx, slots[3])
	DrawGridSlot5(ctx, slots[4])
	DrawGridSlot6(ctx, slots[5])
	DrawGridSlot7(ctx, slots[6])
	DrawGridSlot8(ctx, slots[7])
	DrawGridSlot9(ctx, slots[8])
}
