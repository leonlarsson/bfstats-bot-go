package create

import (
	"github.com/leonlarsson/bfstats-image-gen/canvas"
	"github.com/leonlarsson/bfstats-image-gen/structs"
	core "github.com/tdewolff/canvas"
)

func CreateBF2042Image(data structs.BF2042Data) (*core.Canvas, *core.Context) {
	c, ctx := canvas.CreateStatsCanvasAndContext()

	// Draw some images
	canvas.DrawBackground(ctx, "assets/images/BF2042/BF2042_SOLID_BG_0.png")
	canvas.DrawSkeleton(ctx, canvas.RegularSkeletonType, canvas.RegularStyle)
	canvas.DrawFooterWithText(ctx, "BY MOZZY", "BATTLEFIELDSTATS.COM")
	canvas.DrawGameLogo(ctx, "assets/images/BF2042/BF2042_LOGO_BG.png", canvas.RegularStyle)
	canvas.DrawAvatar(ctx, "assets/images/DefaultGravatar.png")
	canvas.DrawPlatformIcon(ctx, canvas.Platform(data.Platform))

	// Draw some text
	canvas.DrawIdentifier(ctx, data.Identifier)

	canvas.DrawUsername(ctx, data.Username)
	canvas.DrawTimePlayed(ctx, data.TimePlayed)

	canvas.DrawStat1(ctx, data.Stats.Kills.Name, data.Stats.Kills.Value, data.Stats.Kills.Extra)
	canvas.DrawStat2(ctx, data.Stats.Deaths.Name, data.Stats.Deaths.Value, data.Stats.Deaths.Extra)

	canvas.DrawStat3(ctx, data.Stats.Assists.Name, data.Stats.Assists.Value, data.Stats.Assists.Extra)
	canvas.DrawStat4(ctx, data.Stats.Revives.Name, data.Stats.Revives.Value, data.Stats.Revives.Extra)

	canvas.DrawStat5BestClass(ctx, data.Stats.BestClass.Name, data.Stats.BestClass.Value, data.Stats.BestClass.Extra)
	canvas.DrawBestClass(ctx, "assets/images/BF2042/Specialists/Angel.png")

	canvas.DrawStat6(ctx, data.Stats.WlRatio.Name, data.Stats.WlRatio.Value, data.Stats.WlRatio.Extra)

	canvas.DrawRightStat1(ctx, data.Stats.KillsPerMatch.Name, data.Stats.KillsPerMatch.Value, data.Stats.KillsPerMatch.Extra)
	canvas.DrawRightStat2(ctx, data.Stats.KdRatio.Name, data.Stats.KdRatio.Value, data.Stats.KdRatio.Extra)
	canvas.DrawRightStat3(ctx, data.Stats.KillsPerMinute.Name, data.Stats.KillsPerMinute.Value, data.Stats.KillsPerMinute.Extra)
	canvas.DrawRightStat4Rank(ctx, data.Stats.Rank.Name, data.Stats.Rank.Value, data.Stats.Rank.Extra)

	return c, ctx
}
