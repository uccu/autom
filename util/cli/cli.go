package cli

import (
	"autom/conf"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/uccu/go-stringify"
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
						Value:   conf.Base.ConfPath,
						Usage:   "configuration file path",
					},
					&cli.StringFlag{
						Name:    "workdir",
						Aliases: []string{"e"},
						Value:   conf.Base.WorkDir,
						Usage:   "workdir path",
					},
					&cli.StringFlag{
						Name:    "logger",
						Aliases: []string{"l"},
						Value:   conf.Log.Path,
						Usage:   "logger file path",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   conf.Http.Port,
						Usage:   "port number used by the service",
					},
					&cli.BoolFlag{
						Name:    "detach",
						Aliases: []string{"d"},
						Usage:   "run in background",
					},
				},
				Usage: "start the service",
				Action: func(c *cli.Context) error {

					conf.Base.ConfPath = c.String("config")
					conf.Base.WorkDir = c.String("workdir")
					conf.Log.Path = c.String("logger")
					conf.Http.Port = c.Int("port")

					if c.Bool("detach") {
						cmd := exec.Command(
							os.Args[0], "start",
							"-c", conf.Base.ConfPath,
							"-e", conf.Base.WorkDir,
							"-l", conf.Log.Path,
							"-p", stringify.ToString(conf.Http.Port),
						)
						return cmd.Start()
					}

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
			{
				Name:  "version",
				Usage: "print version of autom",
				Action: func(c *cli.Context) error {
					fmt.Println("v0.1.3")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
