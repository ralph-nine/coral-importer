package pipeline

import (
	"runtime"

	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
)

// CollectionWriter will write a document (serialized) out to the collection
// import file.
type CollectionWriter func(collection string, doc easyjson.Marshaler) error

// WritingProcessor is the actual processor that receives a line.
type WritingProcessor func(writer CollectionWriter, input *TaskReaderInput) error

// HandleWritingProcessor will wrap a Processor around a fanning input queue and
// collect the results into an output queue.
func HandleWritingProcessor(in <-chan TaskReaderInput, process WritingProcessor) <-chan TaskWriterOutput {
	out := make(chan TaskWriterOutput)

	writeToOutput := func(collection string, doc easyjson.Marshaler) error {
		bytes, err := easyjson.Marshal(doc)
		if err != nil {
			return errors.Wrap(err, "could not marshal output")
		}

		out <- TaskWriterOutput{
			Collection: collection,
			Document:   bytes,
		}

		return nil
	}

	go func() {
		defer close(out)
		for n := range in {
			if n.Error != nil {
				out <- TaskWriterOutput{
					Error: errors.Wrap(n.Error, "error occurred on stack"),
				}

				return
			}

			if err := process(writeToOutput, &n); err != nil {
				out <- TaskWriterOutput{
					Error: errors.Wrap(err, "error occurred during processing"),
				}

				return
			}
		}
	}()

	return out
}

// FanWritingProcessors will fan the processor across the number of CPU's available.
func FanWritingProcessors(input <-chan TaskReaderInput, process WritingProcessor) []<-chan TaskWriterOutput {
	out := make([]<-chan TaskWriterOutput, runtime.NumCPU())
	for i := range out {
		out[i] = HandleWritingProcessor(input, process)
	}

	return out
}
