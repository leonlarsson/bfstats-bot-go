package canvas

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/leonlarsson/bfstats-go/internal/shared"
	"github.com/leonlarsson/bfstats-go/internal/utils"
	"github.com/tdewolff/canvas"
)

// DrawBackground draws a background image
func DrawBackground(ctx *canvas.Context, filePath string, addDarkeningLayer bool) error {
	bgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return err
	}

	ctx.DrawImage(0, 0, bgImg, 1)

	if addDarkeningLayer {
		ctx.SetFillColor(canvas.RGBA(0, 0, 0, .3))
		rect := canvas.Rectangle(ctx.Width(), ctx.Height())
		ctx.DrawPath(0, 0, rect)
	}

	return nil
}

// DrawSkeleton draws a skeleton background image (the rectangles and scaffolding)
func DrawSkeleton(ctx *canvas.Context, skeletonType shared.SkeletonType, style shared.DrawStyle) error {
	filePath := "assets/images/Skeleton_BGs/Regular.png"

	if skeletonType == shared.GridSkeletonType {
		filePath = "assets/images/Skeleton_BGs/Grid.png"
	}

	if style == shared.DrawnStyle {
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

	ctx.DrawImage(0, 0, bgImg, 1)

	return nil
}

// DrawFooterWithText draws a footer with two texts (left and right).
func DrawFooterWithText(ctx *canvas.Context, leftText string, rightText string, skeletonType shared.SkeletonType) {
	ctx.SetFillColor(canvas.RGBA(32, 32, 32, 0.7))
	leftTextLine := canvas.NewTextLine(robotoFont.Face(PixelsToPoints(20), canvas.RGBA(255, 255, 255, 0.8), canvas.FontMedium), leftText, canvas.Left)
	rightTextLine := canvas.NewTextLine(robotoFont.Face(PixelsToPoints(20), canvas.RGBA(255, 255, 255, 0.8), canvas.FontMedium), rightText, canvas.Right)

	if skeletonType == shared.RegularSkeletonType {
		rect := canvas.Rectangle(700, 25)
		ctx.DrawPath(0, 725, rect)
		ctx.DrawText(35, 745, leftTextLine)
		ctx.DrawText(665, 745, rightTextLine)
	} else {
		rect := canvas.Rectangle(1200, 25)
		ctx.DrawPath(0, 725, rect)
		ctx.DrawText(35, 745, leftTextLine)
		ctx.DrawText(1165, 745, rightTextLine)
	}
}

// DrawGameLogo draws a game logo (game logo files are expected to already be placed in the correct position)
func DrawGameLogo(ctx *canvas.Context, filePath string, style shared.DrawStyle) error {
	if style == shared.DrawnStyle {
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

	ctx.DrawImage(0, 0, bgImg, 1)

	return nil
}

// DrawPlatformIcon draws a platform icon. Dirty way to make drawn optional lol
func DrawPlatformIcon(ctx *canvas.Context, requestedPlatform shared.Platform, skeletonType shared.SkeletonType, drawn bool) error {
	platforms := map[int]string{
		int(shared.PlatformPC):   "assets/images/Platform Icons/PC.png",
		int(shared.PlatformXBOX): "assets/images/Platform Icons/Xbox.png",
		int(shared.PlatformPS):   "assets/images/Platform Icons/PS.png",
	}

	platformImagePath, ok := platforms[int(requestedPlatform)]
	if !ok {
		return fmt.Errorf("platform %d not found", requestedPlatform)
	}

	// If drawn is true use "assets/images/Platform Icons/{PLATFORM}_Drawn.png" instead
	if drawn {
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
	if skeletonType == shared.RegularSkeletonType {
		ctx.FitImage(platformImage, canvas.Rect{X: 1088, Y: 73, W: 75, H: 75}, canvas.ImageContain)
	} else {
		ctx.FitImage(platformImage, canvas.Rect{X: 1125, Y: 36, W: 39, H: 39}, canvas.ImageContain)
	}

	return nil
}

// DrawAvatar draws an avatar image. Local path or URL is accepted.
func DrawAvatar(ctx *canvas.Context, filePath string) error {

	var avatarImg image.Image

	// If filePath is a URL, download the image. Otherwise, open the file
	if strings.HasPrefix(filePath, "http") {
		res, err := http.Get(filePath)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// Decode (works for PNG and JPEG image formats)
		// If webp is needed in the future, webp.Decode() can be used
		avatarImg, _, err = image.Decode(res.Body)

		if err != nil {
			return err
		}

	} else {
		avatarImgFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer avatarImgFile.Close()

		// Assume PNG because I only have PNGs
		avatarImg, err = png.Decode(avatarImgFile)
		if err != nil {
			return err
		}
	}

	// TODO: Figure out how to mask the image into a circle (skill issue). Ref: https://github.com/tdewolff/canvas/issues/232
	ctx.FitImage(avatarImg, canvas.Rect{X: 850, Y: 11, W: 200, H: 200}, canvas.ImageContain)

	return nil
}

// DrawBestClassImage draws the best class image in stat slot 5
func DrawBestClassImage(ctx *canvas.Context, game string, className string) {
	game = strings.ToUpper(game)

	var filePath string

	isBF2042BaseClass := utils.IsBaseBF2042Class(game, className)

	if isBF2042BaseClass {
		filePath = fmt.Sprintf("assets/images/BF2042/Specialists/%s.png", className)
	} else {
		filePath = fmt.Sprintf("assets/images/%s/Classes/%s.png", game, className)
	}

	bestClassFile, err := os.Open(filePath)
	if err != nil {
		// If the file is not found, draw nothing for now
		return
	}
	defer bestClassFile.Close()

	bestClassImg, err := png.Decode(bestClassFile)
	if err != nil {
		// If the file could not be decoded, draw nothing for now
		return
	}

	// If the best class is a BF2042 base class, draw the image with a specific size and position
	if isBF2042BaseClass {
		ctx.FitImage(bestClassImg, canvas.Rect{X: 60, Y: 565, W: 50, H: 50}, canvas.ImageContain)
	} else {
		// TODO: Verify for games that have class icons instead
		ctx.FitImage(bestClassImg, canvas.Rect{X: 40, Y: 552, W: 70, H: 70}, canvas.ImageContain)
	}
}
