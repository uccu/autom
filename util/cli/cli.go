package cli

import (
	"autom/conf"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run() {

	app := &cli.App{
		Name:  "autom",
		Usage: "Listening for git push times to automatically deploy docker servers.",
		Commands: []*cli.Command{
			{
				Name: "start",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   conf.ConfPath,
						Usage:   "configuration file path",
					},
					&cli.StringFlag{
						Name:    "logger",
						Aliases: []string{"l"},
						Value:   conf.Log.Path,
						Usage:   "logger file path",
					},
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   conf.Http.Port,
						Usage:   "port number used by the service",
					},
				},
				Usage: "start the service",
				Action: func(c *cli.Context) error {

					conf.ConfPath = c.String("config")
					conf.Base.ConfPath = c.String("config")
					conf.Log.Path = c.String("logger")
					conf.Http.Port = c.String("port")

					return serverStart()
				},
			},
			{
				Name:  "stop",
				Usage: "stop the service",
				Action: func(c *cli.Context) error {
					return serverStop()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
