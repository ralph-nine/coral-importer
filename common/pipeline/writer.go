package pipeline

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// NewFileWriter will write the outputs out based on the output.
func NewFileWriter(folder string, input <-chan TaskOutput) error {
	writers := make(map[string]*bufio.Writer)

	for task := range input {
		if task.Error != nil {
			return errors.Wrap(task.Error, "error occurred on stack")
		}

		writer, ok := writers[task.Collection]
		if !ok {
			// Ensure that the folder exists.
			if _, err := os.Stat(folder); os.IsNotExist(err) {
				if err := os.Mkdir(folder, 0755); err != nil {
					return errors.Wrap(err, "can't make output directory")
				}
			}

			// Create the file to write to.
			f, err := os.Create(filepath.Join(folder, fmt.Sprintf("%s.json", task.Collection)))
			if err != nil {
				return errors.Wrap(err, "could not create file")
			}
			defer f.Close()

			// Wrap this file in a buffered writer.
			writer = bufio.NewWriter(f)
			defer writer.Flush()

			// Link the writer to the map of writers.
			writers[task.Collection] = writer
		}

		// Write the document out.
		writer.Write(task.Document)
		writer.WriteString("\n")
	}

	return nil
}
