package pipeline

import (
	"sync"

	"github.com/pkg/errors"
)

type TaskAggregatorOutput struct {
	Error      error
	Key        string
	Value      string
	Collection string
}

// MergeTaskAggregatorOutputPipelines will collect all results from the input channels
// and output them on a single channel.
func MergeTaskAggregatorOutputPipelines(cs []<-chan TaskAggregatorOutput) <-chan TaskAggregatorOutput {
	var wg sync.WaitGroup
	out := make(chan TaskAggregatorOutput)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan TaskAggregatorOutput) {
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

func NewMapAggregator(input <-chan TaskAggregatorOutput) (map[string]map[string][]string, error) {
	out := make(map[string]map[string][]string)

	for task := range input {
		if task.Error != nil {
			return nil, errors.Wrap(task.Error, "error occurred on stack")
		}

		// Get the results map to add to the collectio results.
		if _, ok := out[task.Collection]; !ok {
			out[task.Collection] = make(map[string][]string)
		}

		// Ensure that the specific key we're adding to already exists.
		if _, ok := out[task.Collection][task.Key]; !ok {
			out[task.Collection][task.Key] = []string{}
		}

		// Push the value onto the map.
		if task.Value != "" {
			out[task.Collection][task.Key] = append(out[task.Collection][task.Key], task.Value)
		}
	}

	return out, nil
}
