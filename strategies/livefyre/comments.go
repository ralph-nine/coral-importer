package livefyre

import (
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common/pipeline"
)

// Time is the time.Time representation that LiveFyre uses.
type Time struct {
	time.Time
}

// UnmarshalJSON is the custom unmarshaler for the input JSON from LiveFyre.
func (t *Time) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02T15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}

	t.Time = tt

	return nil
}

func ConfigureLogger(c *cli.Context) {
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
}

// Comments will handle a data import task for importing comments into Coral
// from a LiveFyre export.
func Comments(c *cli.Context) error {
	ConfigureLogger(c)

	tenantID := c.String("tenantID")
	input := c.String("input")
	output := c.String("output")
	started := time.Now()

	logrus.Info("started")

	// Create the file reader.
	reader := pipeline.NewFileReader(input)

	// Create the processor that will write these entries out.
	if err := pipeline.NewFileWriter(output, pipeline.MergeTaskOutputPipelines(pipeline.WrapProcessors(reader, runtime.NumCPU(), Process(tenantID)))); err != nil {
		logrus.WithError(err).WithField("took", time.Now().Sub(started).String()).Error("finished processing")
		return err
	}

	logrus.WithField("took", time.Now().Sub(started).String()).Info("finished processing")

	return nil
}
