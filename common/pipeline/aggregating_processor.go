package pipeline

import (
	"runtime"

	"github.com/pkg/errors"
)

type AggregationWriter func(collection, key, value string)

type AggregatingProcessor func(writer AggregationWriter, input *TaskReaderInput) error

func HandleAggregatingProcessor(in <-chan TaskReaderInput, process AggregatingProcessor) <-chan TaskAggregatorOutput {
	out := make(chan TaskAggregatorOutput)

	writeToOutput := func(collection, key, value string) {
		out <- TaskAggregatorOutput{
			Collection: collection,
			Key:        key,
			Value:      value,
		}
	}

	go func() {
		defer close(out)
		for n := range in {
			if n.Error != nil {
				out <- TaskAggregatorOutput{
					Error: errors.Wrap(n.Error, "error occurred on stack"),
				}
				return
			}

			if err := process(writeToOutput, &n); err != nil {
				out <- TaskAggregatorOutput{
					Error: errors.Wrap(err, "error occurred during processing"),
				}
				return
			}
		}
	}()

	return out
}

// FanAggregatingProcessor will fan the processor across the number of CPU's available.
func FanAggregatingProcessor(input <-chan TaskReaderInput, process AggregatingProcessor) []<-chan TaskAggregatorOutput {
	out := make([]<-chan TaskAggregatorOutput, runtime.NumCPU())
	for i := range out {
		out[i] = HandleAggregatingProcessor(input, process)
	}
	return out
}
