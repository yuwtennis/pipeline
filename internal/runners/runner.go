package runners

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/dataflow"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
)

type Runner interface {
	Execute(ctx context.Context, pipeline *beam.Pipeline) (beam.PipelineResult, error)
}

type Direct struct{}
type DataFlow struct{}

// Init initializes Pipeline and Scope
func Init() (pipeline *beam.Pipeline, scope beam.Scope) {
	beam.Init()

	pipeline = beam.NewPipeline()
	scope = pipeline.Root()

	return
}

// Execute executes the Pipeline for Direct runner
func (d Direct) Execute(context context.Context, pipeline *beam.Pipeline) (result beam.PipelineResult, err error) {
	result, err = direct.Execute(context, pipeline)

	return
}

// Execute executes the Pipeline for DataFlow runner
func (d DataFlow) Execute(context context.Context, pipeline *beam.Pipeline) (result beam.PipelineResult, err error) {
	result, err = dataflow.Execute(context, pipeline)

	return
}
