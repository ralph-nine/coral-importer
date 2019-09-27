package common

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// ConfigureLogger will configure the global logger based on global flags.
func ConfigureLogger(c *cli.Context) error {
	quiet := c.GlobalBool("quiet")
	json := c.GlobalBool("json")

	if quiet {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if json {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	return nil
}
