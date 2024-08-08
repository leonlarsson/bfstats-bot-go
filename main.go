package main

import (
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/structs"
	"github.com/tdewolff/canvas/renderers"
)

func main() {
	// router := http.NewServeMux()
	// router.HandleFunc("/image/bf2042", handlers.BF2042Handler)
	// http.ListenAndServe(":8080", router)

	// c, _ := create.CreateBF2042OverviewImage(structs.BF2042OverviewData{
	// 	BaseData: structs.BaseData{
	// 		Username:   "MozzyFX",
	// 		Identifier: "FECbLioP0ywuiztPUP",
	// 		Platform:   0,
	// 		Avatar:     "assets/images/DefaultGravatar.png",
	// 		Meta: structs.Meta{
	// 			Game:    "Battlefield 2042",
	// 			Segment: "Overview",
	// 		},
	// 	},
	// 	Stats: structs.BF2042OverviewStats{
	// 		TimePlayed: structs.Stat{
	// 			Name:  "Time Played:",
	// 			Value: "150 hours",
	// 		},
	// 		Kills: structs.Stat{
	// 			Name:  "Kills:",
	// 			Value: "13,637",
	// 			Extra: "Top 13%",
	// 		},
	// 		Deaths: structs.Stat{
	// 			Name:  "Deaths:",
	// 			Value: "3,254",
	// 			Extra: "Top 29%",
	// 		},
	// 		Assists: structs.Stat{
	// 			Name:  "Assists:",
	// 			Value: "9,158",
	// 			Extra: "Top 12%",
	// 		},
	// 		Revives: structs.Stat{
	// 			Name:  "Revives:",
	// 			Value: "705",
	// 			Extra: "Top 29%",
	// 		},
	// 		BestClass: structs.Stat{
	// 			Name:  "Best Class:",
	// 			Value: "Angel",
	// 			Extra: "2,813 kills | 15 hours",
	// 		},
	// 		WlRatio: structs.Stat{
	// 			Name:  "W/L Ratio:",
	// 			Value: "61.8%",
	// 			Extra: "Top 13%",
	// 		},
	// 		KillsPerMatch: structs.Stat{
	// 			Name:  "Kills/Match:",
	// 			Value: "32.55",
	// 			Extra: "Top 7%",
	// 		},
	// 		KdRatio: structs.Stat{
	// 			Name:  "K/D Ratio:",
	// 			Value: "4.19 (3.4)",
	// 			Extra: "Top 3.3%",
	// 		},
	// 		KillsPerMinute: structs.Stat{
	// 			Name:  "Kills/Minute:",
	// 			Value: "1.51",
	// 			Extra: "Top 10%",
	// 		},
	// 		Rank: structs.RankStat{
	// 			Name:    "Rank 114 (S015)",
	// 			Value:   "96% to next rank",
	// 			Extra:   "XP: 7,586,196",
	// 			RankInt: 114,
	// 		},
	// 	},
	// }, shared.SolidBackground)

	c, _ := create.CreateBF2042WeaponsImage(structs.BF2042WeaponsData{
		BaseData: structs.BaseData{
			Username:   "MozzyFX",
			Identifier: "FECbLioP0ywuiztPUP",
			Platform:   0,
			Meta: structs.Meta{
				Game:    "Battlefield 2042",
				Segment: "Overview",
			},
		},
		Weapons: []structs.Stat{
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "16.2% accuracy | 2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
			{
				Name:  "M5A3",
				Value: "1,234 kills",
				Extra: "2.3 KPM",
			},
		},
	}, shared.SolidBackground)

	// Save the image
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}
}
