package livefyre

import (
	"strings"
	"time"

	"github.com/pkg/errors"
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

// Import will handle a data import task for importing comments into Coral
// from a LiveFyre export.
func Import(c *cli.Context) error {
	// tenantID is the ID of the Tenant that we are importing these documents
	// for.
	tenantID := c.String("tenantID")

	// commentsFileName is the name of the file that that contains the comment
	// exports from LiveFyre.
	commentsFileName := c.String("comments")

	// usersFileName is the name of the file that that contains the comment
	// exports from LiveFyre.
	usersFileName := c.String("users")

	// folder is the name of the folder where we are placing our outputted dumps
	// ready for MongoDB import.
	folder := c.String("output")

	// Mark when we started.
	started := time.Now()
	logrus.Info("started")

	// Process the users. The output users map will map any user's ID to their
	// mapped user ID.
	users, err := HandleUsers(tenantID, folder, usersFileName)
	if err != nil {
		logrus.WithError(err).Error("could not handle")
		return err
	}

	// Create the processor that will write these entries out.
	if err := pipeline.NewFileWriter(
		folder,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewFileReader(commentsFileName),
				ProcessComments(tenantID, users),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process comments and stories for writing")
		return err
	}

	// Mark when we finished.
	finished := time.Now()
	logrus.WithField("took", finished.Sub(started).String()).Info("finished processing")

	return nil
}

func HandleUsers(tenantID, folder, usersFileName string) (map[string]string, error) {

	// Process the users file first because we need to de-duplicate users as
	// they are parsed because LiveFyre did not lowercase email addresses,
	// causing multiple users to be created for each email address variation.
	users, err := pipeline.NewMapAggregator(
		pipeline.MergeTaskAggregatorOutputPipelines(
			pipeline.FanAggregatingProcessor(
				pipeline.NewFileReader(usersFileName),
				ProcessUsersMap(),
			),
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not aggregate users")
	}

	logrus.WithField("users", len(users["id"])).Info("loaded users")

	if err := pipeline.NewFileWriter(
		folder,
		ProcessUsers(tenantID, users),
	); err != nil {
		return nil, errors.Wrap(err, "could not write out users")
	}

	logrus.WithField("users", len(users["id"])).Info("wrote users")

	// Genreate the users association from the id map.
	out := make(map[string]string)
	for _, ids := range users["id"] {
		for _, id := range ids {
			out[id] = ids[0]
		}
	}

	return out, nil
}
