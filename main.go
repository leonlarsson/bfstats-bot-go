package main

import (
	"net/http"

	"github.com/leonlarsson/bfstats-image-gen/handlers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/image/bf2042", handlers.BF2042Handler)
	http.ListenAndServe(":8080", router)
}
