package pipeline

// Process will wrap the processor and input with a series of pipelien managers
// that will fan out the work and write the results out to files.
func Process(folder string, reader <-chan TaskInput, processor Processor) error {
	return NewFileWriter(folder, MergeTaskOutputPipelines(WrapProcessors(reader, processor)))
}
