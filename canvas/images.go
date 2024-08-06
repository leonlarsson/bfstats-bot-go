package canvas

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/tdewolff/canvas"
)

type DrawStyle int

const (
	RegularStyle DrawStyle = iota
	DrawnStyle
)

// DrawBackground draws a background image
func DrawBackground(ctx *canvas.Context, filePath string) error {
	bgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return err
	}

	ctx.DrawImage(0, 0, bgImg, canvas.DPMM(1))

	return nil
}

type SkeletonType int

const (
	RegularSkeletonType SkeletonType = iota
	GridSkeletonType
)

// DrawSkeleton draws a skeleton background image (the rectangles and scaffolding)
func DrawSkeleton(ctx *canvas.Context, skeletonType SkeletonType, style DrawStyle) error {
	filePath := "assets/images/Skeleton_BGs/Regular.png"

	if skeletonType == GridSkeletonType {
		filePath = "assets/images/Skeleton_BGs/Grid.png"
	}

	if style == DrawnStyle {
		filePath = strings.ReplaceAll(filePath, ".png", "_Drawn.png")
	}

	bgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return err
	}

	ctx.DrawImage(0, 0, bgImg, canvas.DPMM(1))

	return nil
}

func DrawFooterWithText(ctx *canvas.Context, leftText string, rightText string) {
	ctx.SetFillColor(canvas.RGBA(32, 32, 32, 0.7))
	rect := canvas.Rectangle(700, 25)
	ctx.DrawPath(0, 725, rect)

	ctx.DrawText(35, 745, canvas.NewTextLine(robotoFont.Face(PixelsToPoints(20), canvas.RGBA(255, 255, 255, 0.8), canvas.FontMedium), leftText, canvas.Left))
	ctx.DrawText(665, 745, canvas.NewTextLine(robotoFont.Face(PixelsToPoints(20), canvas.RGBA(255, 255, 255, 0.8), canvas.FontMedium), rightText, canvas.Right))
}

// DrawGameLogo draws a game logo (game logo files are expected to already be placed in the correct position)
func DrawGameLogo(ctx *canvas.Context, filePath string, style DrawStyle) error {
	if style == DrawnStyle {
		filePath = strings.ReplaceAll(filePath, "_LOGO_BG", "_LOGO_WORDART_BG")
	}

	bgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return err
	}

	ctx.DrawImage(0, 0, bgImg, canvas.DPMM(1))

	return nil
}

type Platform int

const (
	PlatformPC Platform = iota
	PlatformPS
	PlatformXBOX
)

// DrawPlatformIcon draws a platform icon. Dirty way to make drawn optional lol
func DrawPlatformIcon(ctx *canvas.Context, requestedPlatform Platform, drawn ...bool) error {
	platforms := map[int]string{
		int(PlatformPC):   "assets/images/Platform Icons/PC.png",
		int(PlatformXBOX): "assets/images/Platform Icons/Xbox.png",
		int(PlatformPS):   "assets/images/Platform Icons/PS.png",
	}

	platformImagePath, ok := platforms[int(requestedPlatform)]
	if !ok {
		return fmt.Errorf("platform %d not found", requestedPlatform)
	}

	// If drawn is true use "assets/images/Platform Icons/{PLATFORM}_Drawn.png" instead
	if drawn != nil && drawn[0] {
		platformImagePath = platformImagePath[:len(platformImagePath)-4] + "_Drawn.png"
	}

	platformImageFile, err := os.Open(platformImagePath)
	if err != nil {
		return err
	}
	defer platformImageFile.Close()

	platformImage, err := png.Decode(platformImageFile)
	if err != nil {
		return err
	}

	// Draw the image
	// 1.3 because that looks more fitting
	ctx.DrawImage(1088, 73, platformImage, canvas.DPMM(1.3))

	return nil
}

// DrawAvatar draws an avatar image
func DrawAvatar(ctx *canvas.Context, filePath string) error {
	bgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return err
	}

	// TODO: Figure out how to mask the image into a circle (skill issue)
	circlePath := canvas.Circle(100)
	ctx.DrawPath(950, 111, circlePath)
	ctx.DrawImage(850, 11, bgImg, canvas.DPMM(1))
	ctx.Close()

	return nil
}

// DrawBestClass draws the best class image in stat slot 5
func DrawBestClass(ctx *canvas.Context, filePath string) {
	bestClassFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer bestClassFile.Close()

	bestClassImg, err := png.Decode(bestClassFile)
	if err != nil {
		panic(err)
	}

	// TODO: Values currently hardcoded for BF2042 Specialists. In the future: Support other games (smaller class icons) and cases where no class is found. Ref: https://github.com/leonlarsson/bfstats-bot/blob/main/src/utils/canvasUtils.ts#L225-L229
	// Image size (256px) / 5 is roughly 50px
	ctx.DrawImage(60, 565, bestClassImg, canvas.DPMM(5))
}
