package pipeline

import (
	"sync"

	"github.com/pkg/errors"
)

type TaskSummerOutput struct {
	Error      error
	Collection string
	Key        string
	Value      int
}

func MergeTaskSummerOutputPipelines(cs []<-chan TaskSummerOutput) <-chan TaskSummerOutput {
	var wg sync.WaitGroup
	out := make(chan TaskSummerOutput)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan TaskSummerOutput) {
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

func NewSummer(input <-chan TaskSummerOutput) (map[string]map[string]int, error) {
	out := make(map[string]map[string]int)

	for task := range input {
		if task.Error != nil {
			return nil, errors.Wrap(task.Error, "error occurred on stack")
		}

		// Get the results map to add to the collection results.
		if _, ok := out[task.Collection]; !ok {
			out[task.Collection] = make(map[string]int)
		}

		// Ensure that the specific key we're adding to already exists.
		if _, ok := out[task.Collection][task.Key]; !ok {
			out[task.Collection][task.Key] = 0
		}

		// Push the value onto the map.
		if task.Value != 0 {
			out[task.Collection][task.Key] += task.Value
		}
	}

	return out, nil
}
