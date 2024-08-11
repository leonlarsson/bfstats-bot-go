package api

import (
	"net/http"

	"github.com/leonlarsson/bfstats-bot-go/apihandlers"
)

func Start() {
	r := http.NewServeMux()
	r.HandleFunc("/bf2042/overview", apihandlers.BF2042OverviewHandler)
	r.HandleFunc("/bf2042/weapons", apihandlers.BF2042WeaponsHandler)
	r.HandleFunc("/bf2042/vehicles", apihandlers.BF2042VehiclesHandler)
	http.ListenAndServe(":8080", r)
}
