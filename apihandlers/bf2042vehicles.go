package apihandlers

import (
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/leonlarsson/bfstats-bot-go/canvas"
	"github.com/leonlarsson/bfstats-bot-go/canvasdatashapes"
	create "github.com/leonlarsson/bfstats-bot-go/create/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/shared"
)

// BF2042VehiclesHandler handles the /bf2042/vehicles endpoint
func BF2042VehiclesHandler(w http.ResponseWriter, r *http.Request) {
	var data canvasdatashapes.BF2042VehiclesCanvasData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if we have enough vehicles
	if len(data.Vehicles) < 9 {
		http.Error(w, "Not enough weapons", http.StatusBadRequest)
		return
	}

	c, _ := create.CreateBF2042VehiclesImage(data, shared.SolidBackground)
	w.Header().Set("Content-Type", "image/png")
	img := canvas.CanvasToImage(c)
	png.Encode(w, img)
}
