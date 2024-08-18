package httpbot

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/snowflake/v2"
	"github.com/leonlarsson/bfstats-go/internal/httpbot/commands"
)

func handlePing(event *handler.CommandEvent) error {
	return event.CreateMessage(discord.MessageCreate{
		Content: "pong",
		Flags:   discord.MessageFlagEphemeral,
	})
}

// NOTE: Run ngrok http 80 to expose the local server to the internet. Add {ngrok_url}/interactions/callback to the Discord application's interaction endpoint to use this.
func Start(deployCommands bool) {
	var err error

	r := handler.New()
	r.Use(middleware.Logger)

	r.Group(func(r handler.Router) {
		r.Use(middleware.Print("group1"))
		r.Command("/ping", handlePing)
	})

	client, err := disgo.New(os.Getenv("BOT_TOKEN"),
		bot.WithHTTPServerConfigOpts(os.Getenv("BOT_PUBLIC_KEY"),
			httpserver.WithURL("/interactions/callback"),
			httpserver.WithAddress(":80"),
		),
		bot.WithEventListeners(r),
	)
	if err != nil {
		panic("error while building disgo instance: " + err.Error())
	}

	defer client.Close(context.TODO())

	if deployCommands {
		cmds, err := client.Rest().SetGuildCommands(client.ApplicationID(), snowflake.GetEnv("GUILD_ID"), commands.GetCommands())
		if err != nil {
			panic("error while registering commands: " + err.Error())
		}

		fmt.Printf("registered %d commands\n", len(cmds))
	}

	if err = client.OpenHTTPServer(); err != nil {
		panic("error while starting http server: " + err.Error())
	}

	slog.Info("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
