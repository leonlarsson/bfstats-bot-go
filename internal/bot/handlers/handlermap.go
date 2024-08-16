package handlers

import "github.com/leonlarsson/bfstats-go/internal/bot/handlers/bf2042"

// HandlersMap is a map of command names to their respective handler functions
// Mimics the JS implementation
var HandlersMap = map[string]interface{}{
	"bf2042overview": bf2042.HandleBF2042OverviewCommand,
	"bf2042weapons":  bf2042.HandleBF2042WeaponsCommand,
	"bf2042vehicles": bf2042.HandleBF2042VehiclesCommand,
}
