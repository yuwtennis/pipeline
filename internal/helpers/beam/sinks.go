package beam

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"pipelines/internal/elements"
)

var (
	cfg elasticsearch8.Config = elasticsearch8.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, _ = elasticsearch8.NewClient(cfg)
)

//ToElasticsearchFn indexes entity as elasticsearch document
func ToElasticsearchFn(e elements.Element) *esapi.Response {
	doc := e.ToByte()

	h := sha256.New()
	h.Write([]byte(doc))

	req := esapi.IndexRequest{
		Index:      "realestate",
		DocumentID: fmt.Sprintf("%x", h.Sum(nil)),
		Body:       bytes.NewReader(doc),
		Refresh:    "true",
	}

	resp, _ := req.Do(context.Background(), es)

	return resp
}
