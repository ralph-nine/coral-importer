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
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet",
			Usage: "make output quieter",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "output logs in JSON",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "livefyre",
			Subcommands: []cli.Command{
				{
					Name:   "comments",
					Action: livefyre.Comments,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "tenantID",
							Usage:    "ID of the Tenant to import the Comments for",
							Required: true,
						},
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
