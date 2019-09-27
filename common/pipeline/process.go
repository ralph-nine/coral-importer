package pipeline

import (
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
)

// CollectionWriter will write a document (serialized) out to the collection
// import file.
type CollectionWriter func(collection string, doc easyjson.Marshaler) error

// Processor is the actual processor that recieves a line.
type Processor func(writer CollectionWriter, input TaskInput) error

// WrapProcess will wrap a Processor around a fanning input queue and
// collect the results into an output queue.
func WrapProcess(in <-chan TaskInput, process Processor) <-chan TaskOutput {
	out := make(chan TaskOutput)

	writeToOutput := func(collection string, doc easyjson.Marshaler) error {
		bytes, err := easyjson.Marshal(doc)
		if err != nil {
			return errors.Wrap(err, "could not marshal output")
		}

		out <- TaskOutput{
			Collection: collection,
			Document:   bytes,
		}

		return nil
	}

	go func() {
		defer close(out)
		for n := range in {
			if n.Error != nil {
				out <- TaskOutput{
					Error: errors.Wrap(n.Error, "error occurred on stack"),
				}
				return
			}

			if err := process(writeToOutput, n); err != nil {
				out <- TaskOutput{
					Error: errors.Wrap(err, "error occurred during processing"),
				}
				return
			}
		}
	}()

	return out
}

func WrapProcessors(input <-chan TaskInput, size int, process Processor) []<-chan TaskOutput {
	out := make([]<-chan TaskOutput, size)
	for i := range out {
		out[i] = WrapProcess(input, process)
	}
	return out
}
