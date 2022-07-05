package direct

import "github.com/apache/beam/sdks/v2/go/pkg/beam"

type Direct struct{}

// Init initializes Pipeline and Scope of direct runner
func (d *Direct) Init() (*beam.Pipeline, beam.Scope) {
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	return p, s
}
