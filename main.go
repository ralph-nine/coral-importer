package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/strategies/legacy"
	"gitlab.com/coralproject/coral-importer/strategies/livefyre"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "coral-importer"
	app.Usage = "imports comment exports from other providers into Coral"
	app.Version = fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
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
	app.Before = common.ConfigureLogger
	app.Commands = []cli.Command{
		{
			Name:   "livefyre",
			Action: livefyre.Import,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "tenantID",
					Usage:    "ID of the Tenant to import for",
					Required: true,
				},
				cli.StringFlag{
					Name:     "comments",
					Usage:    "newline separated JSON input file containing comments",
					Required: true,
				},
				cli.StringFlag{
					Name:     "users",
					Usage:    "newline separated JSON input file containing users",
					Required: true,
				},
				cli.StringFlag{
					Name:     "output",
					Usage:    "folder where the outputted mongo files should be placed",
					Required: true,
				},
			},
		},
		{
			Name:   "legacy",
			Action: legacy.Import,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "tenantID",
					Usage:    "ID of the Tenant to import for",
					Required: true,
				},
				cli.StringFlag{
					Name:     "input",
					Usage:    "folder where the output from mongoexport is located, separated into collection named JSON files",
					Required: true,
				},
				cli.StringFlag{
					Name:     "output",
					Usage:    "folder where the outputted mongo files should be placed",
					Required: true,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal()
	}
}
