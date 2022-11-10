package pipelines

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"pipelines/internal/helpers"
	elasticsearchio "pipelines/internal/io"
	"pipelines/internal/runners"
)

type XLangElasticsearch struct {
	// Address of cross language transform sdk aka expansion service
	ExpansionAddr string

	// Array of elasticsearch node addresses
	EsNodeAddrs []string

	// Index name which you are indexing to
	IndexName string

	// Mapping type
	MappingType string
}

func (xle XLangElasticsearch) Process(runner runners.Runner) {
	pipeline, scope := runners.Init()
	pCol := beam.CreateList(scope, []string{"a"})

	elasticsearchio.Write(
		scope,
		xle.ExpansionAddr,
		xle.EsNodeAddrs,
		xle.IndexName,
		xle.MappingType,
		pCol)

	result, err := runner.Execute(
		context.Background(),
		pipeline)
	helpers.Check(err)
	fmt.Println(result)
}
