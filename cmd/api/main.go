package main

import (
	"log"

	"github.com/leonlarsson/bfstats-go/internal/api"
)

func main() {
	log.Println("API: Starting")
	api.Start(":8080")
}
