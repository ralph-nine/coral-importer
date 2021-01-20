package pipeline

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
)

type TaskWriterOutput struct {
	Error      error
	Collection string
	Document   []byte
}

// MergeTaskWriterOutputPipelines will collect all results from the input channels
// and output them on a single channel.
func MergeTaskWriterOutputPipelines(cs []<-chan TaskWriterOutput) <-chan TaskWriterOutput {
	var wg sync.WaitGroup
	out := make(chan TaskWriterOutput)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan TaskWriterOutput) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// NewFileWriter will write the outputs out based on the output.
func NewFileWriter(folder string, input <-chan TaskWriterOutput) error {
	writers := make(map[string]*bufio.Writer)

	for task := range input {
		if task.Error != nil {
			return errors.Wrap(task.Error, "error occurred on stack")
		}

		writer, ok := writers[task.Collection]
		if !ok {
			// Ensure that the folder exists.
			if _, err := os.Stat(folder); os.IsNotExist(err) {
				if err := os.Mkdir(folder, 0o755); err != nil {
					return errors.Wrap(err, "can't make output directory")
				}
			}

			// Create the file to write to.
			f, err := os.Create(filepath.Join(folder, fmt.Sprintf("%s.json", task.Collection)))
			if err != nil {
				return errors.Wrap(err, "could not create file")
			}
			//nolint:staticcheck
			defer f.Close()

			// Wrap this file in a buffered writer.
			writer = bufio.NewWriter(f)
			//nolint:staticcheck
			defer writer.Flush()

			// Link the writer to the map of writers.
			writers[task.Collection] = writer
		}

		// Write the document out.
		if _, err := writer.Write(task.Document); err != nil {
			return errors.Wrap(err, "could not write")
		}

		if _, err := writer.WriteString("\n"); err != nil {
			return errors.Wrap(err, "could not write")
		}
	}

	return nil
}
