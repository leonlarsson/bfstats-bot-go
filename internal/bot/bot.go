package bot

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/snowflake/v2"
	"github.com/leonlarsson/bfstats-go/internal/bot/commands"
)

// NOTE: Run ngrok http 80 to expose the local server to the internet. Add {ngrok_url}/interactions/callback to the Discord application's interaction endpoint to use this.
func Start(deployCommands bool) {
	var err error

	client, err := disgo.New(os.Getenv("BOT_TOKEN"),
		bot.WithHTTPServerConfigOpts(os.Getenv("BOT_PUBLIC_KEY"),
			httpserver.WithURL("/interactions/callback"),
			httpserver.WithAddress(":80"),
		),
		bot.WithEventListeners(Router()),
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
