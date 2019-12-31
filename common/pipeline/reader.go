package pipeline

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

type TaskReaderInput struct {
	Error  error
	Line   int
	Input  string
	Fields []string
}

func NewCSVFileReader(fileName string, fieldsPerRecord int) <-chan TaskReaderInput {
	out := make(chan TaskReaderInput)
	go func() {
		defer close(out)

		// Open that file for reading.
		f, err := os.Open(fileName)
		if err != nil {
			out <- TaskReaderInput{
				Error: errors.Wrap(err, "could not open --input for reading"),
			}
			return
		}
		defer f.Close()

		// Setup the reader.
		r := csv.NewReader(bufio.NewReader(f))
		r.FieldsPerRecord = fieldsPerRecord
		r.TrimLeadingSpace = true

		// Keep track of the processed lines.
		lines := 0

		// Start reading the stories line by line from the file.
		for {
			fields, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}

				out <- TaskReaderInput{
					Error: errors.Wrap(err, "couldn't read the file"),
				}
				return
			}

			// Increment the line count.
			lines++

			// Send the input to a processor.
			out <- TaskReaderInput{
				Line:   lines,
				Fields: fields,
			}
		}
	}()

	return out
}

// NewJSONFileReader will read from a file and emit new TaskInput's from the lines.
func NewJSONFileReader(fileName string) <-chan TaskReaderInput {
	out := make(chan TaskReaderInput)
	go func() {
		defer close(out)

		// Open that file for reading.
		f, err := os.Open(fileName)
		if err != nil {
			out <- TaskReaderInput{
				Error: errors.Wrap(err, "could not open --input for reading"),
			}
			return
		}
		defer f.Close()

		// Setup the scanner.
		r := bufio.NewReader(f)

		// Keep track of the processed lines.
		lines := 0

		// Start reading the stories line by line from the file.
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				out <- TaskReaderInput{
					Error: errors.Wrap(err, "couldn't read the file"),
				}
				return
			}

			// Increment the line count.
			lines++

			// Send the input to a processor.
			out <- TaskReaderInput{
				Line:  lines,
				Input: line,
			}
		}
	}()

	return out
}
