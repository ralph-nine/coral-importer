package pipeline

import (
	"runtime"

	"github.com/pkg/errors"
)

type SummerWriter func(collection, key string, value int)

type SummerProcessor func(writer SummerWriter, input *TaskReaderInput) error

func HandleSummerProcessor(in <-chan TaskReaderInput, process SummerProcessor) <-chan TaskSummerOutput {
	out := make(chan TaskSummerOutput)

	writeToOutput := func(collection, key string, value int) {
		out <- TaskSummerOutput{
			Collection: collection,
			Key:        key,
			Value:      value,
		}
	}

	go func() {
		defer close(out)
		for n := range in {
			if n.Error != nil {
				out <- TaskSummerOutput{
					Error: errors.Wrap(n.Error, "error occured on stack"),
				}
				return
			}

			if err := process(writeToOutput, &n); err != nil {
				out <- TaskSummerOutput{
					Error: errors.Wrap(err, "error occurred during processing"),
				}
				return
			}
		}
	}()

	return out
}

// FanSummerProcessor will fan the processor across the number of CPU's available.
func FanSummerProcessor(input <-chan TaskReaderInput, process SummerProcessor) []<-chan TaskSummerOutput {
	out := make([]<-chan TaskSummerOutput, runtime.NumCPU())
	for i := range out {
		out[i] = HandleSummerProcessor(input, process)
	}
	return out
}
