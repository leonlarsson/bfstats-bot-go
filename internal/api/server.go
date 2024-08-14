package api

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	r := http.NewServeMux()

	for route, handlerFunc := range Routes {
		r.HandleFunc(route, handlerFunc)
	}

	log.Printf("API: Listening on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Printf("API: ListenAndServe error: %v", err)
	}
}
