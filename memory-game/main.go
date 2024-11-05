package main

import (

	"context"
	"fmt"
	"github.com/gospodinzerkalo/memory-game/internal"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
)

var (
	configPath 	= "./"
	tgToken = ""

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config, c",
			Usage:       "path to .env config file",
			Aliases: []string{"c"},
			Destination: &configPath,
		},
	}
)

func main() {

	app := cli.NewApp()
	// app.Flags = flags
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the bot",
			Action: StartBot,
			Flags: flags,
		},
	}
	app.Run(os.Args)

}

func parseEnv() {
	if configPath != "" {
		godotenv.Overload(configPath)
	}
	tgToken = os.Getenv("TG_TOKEN")
	if tgToken == "" {
		panic("Telegram token is required")
	}
}

func StartBot(d *cli.Context) error {
	parseEnv()
	b, err := tb.NewBot(tb.Settings{
		Token:  tgToken,
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err, "DD")
		return err
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	repo := t_bot.PostgreUser(t_bot.PostgreConfig{
		User:     "postgres",
		Password: "postgres",
		Port:     "5432",
		Host:     "db",
	})

	endpoints := t_bot.NewEndpointsFactory(repo, ctx)

	b.Handle("/start", endpoints.Start(b))
	b.Handle(&t_bot.NewGameReplyButton, endpoints.NewGame(b))
	b.Handle(&t_bot.EasyButton, endpoints.StartGameWithEasy(b))
	b.Handle(&t_bot.MediumButton, endpoints.StartGameWithMedium(b))
	b.Handle(&t_bot.HardButton, endpoints.StartGameWithHard(b))
	b.Handle(&t_bot.NoobButton, endpoints.StartGameWithNoob(b))
	b.Handle(&t_bot.GlobalRatingButton, endpoints.GlobalRating(b))
	b.Handle(&t_bot.MyRatingButton, endpoints.MyRating(b))
	b.Handle(tb.OnText, endpoints.Help(b))

	fmt.Println("BOT STARTED!!!")
	b.Start()
	return nil
}
