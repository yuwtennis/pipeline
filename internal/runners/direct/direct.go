package direct

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
)

type Direct struct {
	// Pipeline
	Pipeline *beam.Pipeline

	// Beam Scope
	Scope beam.Scope
}

// Init initializes Pipeline and Scope of direct runner
func (d *Direct) Init() {
	beam.Init()

	d.Pipeline = beam.NewPipeline()
	d.Scope = d.Pipeline.Root()
}

func (d *Direct) Execute(context context.Context, pipeline *beam.Pipeline) {
	direct.Execute(context, pipeline)
}
