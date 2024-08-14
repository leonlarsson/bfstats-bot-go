package api

import (
	"net/http"

	"github.com/leonlarsson/bfstats-go/internal/api/handlers"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/image/bf2042/overview": handlers.BF2042OverviewHandler,
	"/image/bf2042/weapons":  handlers.BF2042WeaponsHandler,
	"/image/bf2042/vehicles": handlers.BF2042VehiclesHandler,
}
