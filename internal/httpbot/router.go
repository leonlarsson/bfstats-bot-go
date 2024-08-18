package httpbot

import (
	"cmp"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/leonlarsson/bfstats-go/internal/httpbot/events"
	"github.com/leonlarsson/bfstats-go/internal/httpbot/handlers/bf2042"
	"github.com/leonlarsson/bfstats-go/internal/httpbot/handlers/misc"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

func Router() *handler.Mux {
	r := handler.New()
	r.Use(middleware.Go)

	r.Autocomplete("/{game}/stats", func(e *handler.AutocompleteEvent) error {
		return events.HandleAutocomplete(e.AutocompleteInteractionCreate)
	})

	r.SlashCommand("/bf2042/stats", func(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
		segment := data.String("segment")
		loc := *localization.CreateLocForLanguage(cmp.Or(data.String("language"), e.Locale().Code(), "en"))

		switch segment {
		case "overview":
			return bf2042.HandleBF2042OverviewCommand(e.ApplicationCommandInteractionCreate, loc)
		case "weapons":
			return bf2042.HandleBF2042WeaponsCommand(e.ApplicationCommandInteractionCreate, loc)
		case "vehicles":
			return bf2042.HandleBF2042VehiclesCommand(e.ApplicationCommandInteractionCreate, loc)
		default:
			return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Segment not implemented").Build())
		}

	})

	r.SlashCommand("/help", func(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
		return misc.HandleHelp(e.ApplicationCommandInteractionCreate)
	})

	r.SlashCommand("/about", func(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
		return misc.HandleAbout(e.ApplicationCommandInteractionCreate)
	})

	r.NotFound(func(e *handler.InteractionEvent) error {
		return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Not found").Build())
	})

	return r
}
