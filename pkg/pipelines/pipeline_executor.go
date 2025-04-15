package pipelines

import (
	"context"
	"sync"
)

// PipelineRunner is a collection of pipelines that can be run in parallel
type PipelineRunner struct {
	Pipelines []Pipeline
}

// NewPipelineRunner creates a PipelineRunner with provided pipelines
func NewPipelineRunner(pipelines []Pipeline) *PipelineRunner {
	return &PipelineRunner{
		Pipelines: pipelines,
	}
}

func (pr *PipelineRunner) Run(ctx context.Context, stop chan<- bool) {
	var wg *sync.WaitGroup

	for _, pipeline := range pr.Pipelines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pipeline.Run(ctx)
		}()
	}
	wg.Wait()
	stop <- true
}
