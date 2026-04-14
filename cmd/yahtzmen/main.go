package main

import (
	"context"
	"log"
	"os"

	"github.com/comsma/yahtzmen/internal/config"
	"github.com/urfave/cli/v3"
)

func main() {
	var cfg *config.Config

	app := &cli.Command{
		Name:  "yahtzmen",
		Usage: "Yahtzmen server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "path to config file",
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			var err error
			cfg, err = config.Load(cmd.String("config"))
			return ctx, err
		},
		Commands: []*cli.Command{
			initCmd(&cfg),
			serveCmd(&cfg),
			seedCmd(&cfg),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
