package handlers

import (
	"encoding/json"
	"image/png"
	"net/http"

	create "github.com/leonlarsson/bfstats-go/create/bf2042"
	"github.com/leonlarsson/bfstats-go/internal/canvas"
	shapes "github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/shared"
)

// BF2042WeaponsHandler handles the /bf2042/weapons endpoint
func BF2042WeaponsHandler(w http.ResponseWriter, r *http.Request) {
	var data shapes.BF2042WeaponsCanvasData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if we have enough weapons
	if len(data.Weapons) < 9 {
		http.Error(w, "Not enough weapons", http.StatusBadRequest)
		return
	}

	c, _ := create.CreateBF2042WeaponsImage(data, shared.SolidBackground)
	w.Header().Set("Content-Type", "image/png")
	img := canvas.CanvasToImage(c)
	png.Encode(w, img)
}
