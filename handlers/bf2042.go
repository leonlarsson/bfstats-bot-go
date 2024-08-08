package handlers

import (
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/create"
	"github.com/leonlarsson/bfstats-bot-go/shared"
	"github.com/leonlarsson/bfstats-bot-go/structs"
)

func BF2042Handler(w http.ResponseWriter, r *http.Request) {
	var data structs.BF2042Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, _ := create.CreateBF2042Image(data, shared.SolidBackground)
	w.Header().Set("Content-Type", "image/png")
	img := canvas.CanvasToImage(c)
	png.Encode(w, img)
}
