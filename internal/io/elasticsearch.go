package elasticsearchio

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/reflectx"
)

const (
	writeURN = "beam:transform:org.apache.beam:elasticsearch_write:v1"
)

type writePayload struct {
	NodeAddrs   []string
	Index       string
	MappingType string
}

// Write is a cross-language PTransform which writes String data to　specified
// Elasticsearch index
//
// Example of Write with required parameters
//
//   expansionAddr := "localhost:1234"
//   nodeAddr := ["localhost:9200"]
//   index := "my-index"
//   type  := "_doc"
func Write(s beam.Scope, addr string, nodeAddrs []string, index string, mappingType string, col beam.PCollection) {
	wpl := writePayload{
		NodeAddrs:   nodeAddrs,
		Index:       index,
		MappingType: mappingType,
	}
	pl := beam.CrossLanguagePayload(wpl)
	namedOutputTypes := map[string]typex.FullType{
		"org.apache.beam.sdk.io.elasticsearch.ElasticsearchIO$Write#0": typex.New(reflectx.ByteSlice),
		"org.apache.beam.sdk.io.elasticsearch.ElasticsearchIO$Write#1": typex.New(reflectx.ByteSlice),
	}
	outputs := beam.CrossLanguage(s, writeURN, pl, addr, beam.UnnamedInput(col), namedOutputTypes)

	if outputs == nil {
		fmt.Println("error")
	}
}
