package canvas

import (
	"fmt"
	"image/png"
	"os"

	"github.com/tdewolff/canvas"
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
		int(PlatformXBOX): "assets/images/Platform Icons/XBOX.png",
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

	// TODO: Figure out how to mask the image into a circle
	circlePath := canvas.Circle(100)
	ctx.DrawPath(950, 111, circlePath)
	ctx.DrawImage(850, 11, bgImg, canvas.DPMM(1))
	ctx.Close()

	return nil
}
