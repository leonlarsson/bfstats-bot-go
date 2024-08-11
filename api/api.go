package api

import (
	"log"
	"net/http"

	"github.com/leonlarsson/bfstats-bot-go/apihandlers"
)

func Start(port string) {
	r := http.NewServeMux()

	handlerMap := map[string]http.HandlerFunc{
		"/bf2042/overview": apihandlers.BF2042OverviewHandler,
		"/bf2042/weapons":  apihandlers.BF2042WeaponsHandler,
		"/bf2042/vehicles": apihandlers.BF2042VehiclesHandler,
	}

	for route, handlerFunc := range handlerMap {
		r.HandleFunc(route, handlerFunc)
	}

	log.Printf("API: Listening on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Printf("API: ListenAndServe error: %v", err)
	}
}
