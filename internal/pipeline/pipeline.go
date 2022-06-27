package pipeline

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"realestatetrans/internal"
)

func Run(events []string) {
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	pCol := beam.Create(s, events)

	// To RealEstate
	// 建築年 & 取引時点 to time object

	err := beamx.Run(context.Background(), p)

	internal.Check(err)
}
