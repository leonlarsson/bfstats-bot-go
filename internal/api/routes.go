package api

import (
	"net/http"

	"github.com/leonlarsson/bfstats-go/internal/api/apihandlers"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/image/bf2042/overview": apihandlers.BF2042OverviewHandler,
	"/image/bf2042/weapons":  apihandlers.BF2042WeaponsHandler,
	"/image/bf2042/vehicles": apihandlers.BF2042VehiclesHandler,
}
