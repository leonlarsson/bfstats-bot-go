package main

import (
	"github.com/leonlarsson/bfstats-bot-go/internal/api"
)

func main() {
	api.StartServer(":8080")
}
