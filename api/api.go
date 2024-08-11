package api

import (
	"log"
	"net/http"

	"github.com/leonlarsson/bfstats-bot-go/apihandlers"
)

func Start(port string) {
	r := http.NewServeMux()
	r.HandleFunc("/bf2042/overview", apihandlers.BF2042OverviewHandler)
	r.HandleFunc("/bf2042/weapons", apihandlers.BF2042WeaponsHandler)
	r.HandleFunc("/bf2042/vehicles", apihandlers.BF2042VehiclesHandler)
	log.Printf("API: Listening on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Printf("API: ListenAndServe error: %v", err)
	}
}
