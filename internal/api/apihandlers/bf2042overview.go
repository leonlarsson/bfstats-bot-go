package apihandlers

import (
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/leonlarsson/bfstats-go/internal/canvas"
	"github.com/leonlarsson/bfstats-go/internal/canvas/shapes"
	"github.com/leonlarsson/bfstats-go/internal/createcanvas/bf2042"
	"github.com/leonlarsson/bfstats-go/internal/shared"
)

// BF2042OverviewHandler handles the /bf2042/overview endpoint
func BF2042OverviewHandler(w http.ResponseWriter, r *http.Request) {
	var data shapes.GenericRegularData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, _ := bf2042.CreateBF2042OverviewImage(data, shared.SolidBackground)
	w.Header().Set("Content-Type", "image/png")
	img := canvas.CanvasToImage(c)
	png.Encode(w, img)
}
