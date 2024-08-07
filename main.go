package main

import (
	"github.com/leonlarsson/bfstats-image-gen/create"
	"github.com/leonlarsson/bfstats-image-gen/shared"
	"github.com/leonlarsson/bfstats-image-gen/structs"
	"github.com/tdewolff/canvas/renderers"
)

func main() {
	// router := http.NewServeMux()
	// router.HandleFunc("/image/bf2042", handlers.BF2042Handler)
	// http.ListenAndServe(":8080", router)

	// Test data
	data := structs.BF2042Data{
		BaseData: structs.BaseData{
			Username:   "MozzyFX",
			Identifier: "FECbLioP0ywuiztPUP",
			TimePlayed: "150 hours",
			Platform:   0,
		},
		Stats: structs.BF2042Stats{
			Kills: structs.Stat{
				Name:  "Kills:",
				Value: "13,637",
				Extra: "Top 13%",
			},
		},
	}

	// Create image with the data
	c, _ := create.CreateBF2042Image(data, shared.SolidBackground)

	// Save the image
	if err := renderers.Write("render.png", c); err != nil {
		panic(err)
	}
}
