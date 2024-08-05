package main

import (
	"github.com/leonlarsson/go-image-gen/canvas"
	core "github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func main() {

	c, ctx := canvas.CreateCanvasAndContext(1200, 750)

	// Set the coordinate system to match what I am used to: (0, 0) in the top left corner
	ctx.SetCoordSystem(core.CartesianIV)

	// Draw some images
	canvas.DrawBackground(ctx, "assets/images/BF2042/BF2042_IMAGE_BG_0.png")
	canvas.DrawBackground(ctx, "assets/images/Skeleton_BGs/Regular.png")
	canvas.DrawBackground(ctx, "assets/images/BF2042/BF2042_LOGO_BG.png")
	canvas.DrawAvatar(ctx, "assets/images/DefaultGravatar.png")
	canvas.DrawPlatformIcon(ctx, canvas.PlatformPS)

	// Draw some text
	canvas.DrawIdentifier(ctx, "zPptelDNG1uE")

	canvas.DrawStat1(ctx, "Kills:", "13,637", "Top 13%")
	canvas.DrawStat2(ctx, "Deaths:", "3,254", "")

	canvas.DrawStat3(ctx, "Assists:", "9,158", "Top 10%")
	canvas.DrawStat4(ctx, "Revives:", "705", "Top 29%")

	canvas.DrawStat5(ctx, "Best Class:", "Angel", "2,818 kills | 15 hours")
	canvas.DrawStat6(ctx, "W/L Ratio:", "61.8%", "Top 13%")

	canvas.DrawRightStat1(ctx, "Kills/Match:", "32.55", "Top 7%")
	canvas.DrawRightStat2(ctx, "K/D Ratio:", "4.19 (3.4)", "Top 3.3%")
	canvas.DrawRightStat3(ctx, "Kills Per Minute:", "1.51", "Top 10%")
	canvas.DrawRightStat4Rank(ctx, "Rank 114 (S015)", "96% to next rank", "XP: 7,586,196")

	// Save the image
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}
}
