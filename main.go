package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/strategies/livefyre"
)

func main() {
	app := cli.NewApp()
	app.Name = "coral-importer"
	app.Usage = "imports comment exports from other providers into Coral"
	app.Commands = []cli.Command{
		{
			Name: "livefyre",
			Subcommands: []cli.Command{
				{
					Name:   "comments",
					Action: livefyre.Comments,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "input",
							Usage:    "newline seperated JSON input file containing comments",
							Required: true,
						},
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("could not setup the CLI")
	}
}
