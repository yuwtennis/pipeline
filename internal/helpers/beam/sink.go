package beam

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

//ToElasticsearch indexes entity as elasticsearch document
func ToElasticsearch(pCol beam.PCollection) {
	es8, _ := elasticsearch8.NewDefaultClient()

}
