package livefyre

import (
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

// Comments will handle a data import task for importing comments into Coral
// from a LiveFyre export.
func Comments(c *cli.Context) error {
	// tenantID is the ID of the Tenant that we are importing these documents
	// for.
	tenantID := c.String("tenantID")

	// Create the processor that will perform the transforming.
	processor := Process(tenantID)

	// fileName is the name of the file that we are reading in to import.
	fileName := c.String("input")

	// folder is the name of the folder where we are placing our outputted dumps
	// ready for MongoDB import.
	folder := c.String("output")

	// Mark when we started.
	started := time.Now()
	logrus.Info("started")

	// Create the processor that will write these entries out.
	if err := pipeline.Process(folder, pipeline.NewFileReader(fileName), processor); err != nil {
		logrus.WithError(err).Error("could not process entities")
		return err
	}

	// Mark when we finished.
	finished := time.Now()
	logrus.WithField("took", finished.Sub(started).String()).Info("finished processing")

	return nil
}
