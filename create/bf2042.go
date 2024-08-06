package create

import (
	"github.com/leonlarsson/bfstats-image-gen/canvas"
	"github.com/leonlarsson/bfstats-image-gen/structs"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042Image(data structs.BaseData) (*core.Canvas, *core.Context) {
	c, ctx := canvas.CreateCanvasAndContext(1200, 750)

	// Set the coordinate system to match what I am used to: (0, 0) in the top left corner
	ctx.SetCoordSystem(core.CartesianIV)

	// Draw some images
	canvas.DrawBackground(ctx, "assets/images/BF2042/BF2042_SOLID_BG_0.png")
	canvas.DrawSkeleton(ctx, canvas.RegularSkeletonType, canvas.RegularStyle)
	canvas.DrawFooterWithText(ctx, "BY MOZZY", "BATTLEFIELDSTATS.COM")
	canvas.DrawGameLogo(ctx, "assets/images/BF2042/BF2042_LOGO_BG.png", canvas.RegularStyle)
	canvas.DrawAvatar(ctx, "assets/images/DefaultGravatar.png")
	canvas.DrawPlatformIcon(ctx, canvas.Platform(data.Platform))

	// Draw some text
	canvas.DrawIdentifier(ctx, "zPptelDNG1uE")

	canvas.DrawUsername(ctx, data.Username)
	canvas.DrawTimePlayed(ctx, data.TimePlayed)

	slot1 := getStatForSlot(data, 1)
	canvas.DrawStat1(ctx, slot1.Name, slot1.Value, slot1.Extra)
	slot2 := getStatForSlot(data, 2)
	canvas.DrawStat2(ctx, slot2.Name, slot2.Value, slot2.Extra)

	slot3 := getStatForSlot(data, 3)
	canvas.DrawStat3(ctx, slot3.Name, slot3.Value, slot3.Extra)
	slot4 := getStatForSlot(data, 4)
	canvas.DrawStat4(ctx, slot4.Name, slot4.Value, slot4.Extra)

	canvas.DrawStat5BestClass(ctx, "Best Class:", "Angel", "2,818 kills | 15 hours")
	canvas.DrawBestClass(ctx, "assets/images/BF2042/Specialists/Angel.png")

	slot6 := getStatForSlot(data, 6)
	canvas.DrawStat6(ctx, slot6.Name, slot6.Value, slot6.Extra)

	canvas.DrawRightStat1(ctx, "Kills/Match:", "32.55", "Top 7%")
	canvas.DrawRightStat2(ctx, "K/D Ratio:", "4.19 (3.4)", "Top 3.3%")
	canvas.DrawRightStat3(ctx, "Kills Per Minute:", "1.51", "Top 10%")
	canvas.DrawRightStat4Rank(ctx, "Rank 114 (S015)", "96% to next rank", "XP: 7,586,196")

	return c, ctx
}

func getStatForSlot(data structs.BaseData, slot uint) *structs.Stat {
	for _, stat := range data.Stats {
		if stat.Slot == slot {
			return &stat
		}
	}
	return nil
}
