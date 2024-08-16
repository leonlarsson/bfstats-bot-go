package handlers

import commandhandlers "github.com/leonlarsson/bfstats-go/internal/bot/handlers/bf2042"

// HandlersMap is a map of command names to their respective handler functions
// Mimics the JS implementation
var HandlersMap = map[string]interface{}{
	"bf2042overview": commandhandlers.HandleBF2042OverviewCommand,
	"bf2042weapons":  commandhandlers.HandleBF2042WeaponsCommand,
	"bf2042vehicles": commandhandlers.HandleBF2042VehiclesCommand,
}
