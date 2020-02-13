package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/strategies/csv"
	"gitlab.com/coralproject/coral-importer/strategies/legacy"
	"gitlab.com/coralproject/coral-importer/strategies/livefyre"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const (
	// CurrentMigrationVersion is the version representing the most recent migration
	// that this strategy is designed to handle. This should be updated as revisions
	// are applied to this strategy for future versions.
	CurrentMigrationVersion int64 = 1580404849316
)

func main() {
	app := cli.NewApp()
	app.Name = "coral-importer"
	app.Usage = "imports comment exports from other providers into Coral"
	app.Version = fmt.Sprintf("%v, commit %v, built at %v against migration %d", version, commit, date, CurrentMigrationVersion)
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet",
			Usage: "make output quieter",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "output logs in JSON",
		},
		cli.Int64Flag{
			Name:     "migrationID",
			Usage:    "ID of the most recent migration associated with your installation",
			Required: true,
		},
		cli.BoolFlag{
			Name:  "forceSkipMigrationCheck",
			Usage: "used to skip the migration version check",
		},
	}
	app.Before = func(c *cli.Context) error {
		// Configure the logger.
		if err := common.ConfigureLogger(c); err != nil {
			return err
		}

		// Check that the imported needs updating.
		if c.GlobalBool("forceSkipMigrationCheck") {
			logrus.Warn("skipping migration check")
		} else if c.GlobalInt64("migrationID") != CurrentMigrationVersion {
			logrus.WithFields(logrus.Fields{
				"migrationID":             c.GlobalInt("migrationID"),
				"currentMigrationVersion": CurrentMigrationVersion,
			}).Fatal("migration version mismatch, update importer to support new migrations or skip with --forceSkipMigrationCheck")
		}

		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:   "csv",
			Usage:  "a migrator designed to migrate data from the standardized CSV format",
			Action: csv.Import,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "tenantID",
					Usage:    "ID of the Tenant to import for",
					Required: true,
				},
				cli.StringFlag{
					Name:     "input",
					Usage:    "folder where the CSV input files are located",
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
			Name:   "livefyre",
			Usage:  "a migrator designed to migrate data from the LiveFyre platform",
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
			Usage:  "a migrator designed to import data from previous versions of Coral",
			Action: legacy.Import,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "tenantID",
					Usage:    "ID of the Tenant to import for",
					Required: true,
				},
				cli.StringFlag{
					Name:  "preferredPerspectiveModel",
					Usage: "the preferred model to use for copying over toxicity scores",
					Value: "SEVERE_TOXICITY",
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
