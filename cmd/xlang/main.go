package main

import (
	"pipelines/internal/pipelines"
	"pipelines/internal/runners"
)

func main() {
	p := pipelines.XLangElasticsearch{
		ExpansionAddr: "localhost:12345",
		EsNodeAddrs:   []string{"192.168.1.2:9200"},
		IndexName:     "MyIndex",
		MappingType:   "_doc",
	}

	var runner runners.Runner
	runner = runners.DataFlow{}

	p.Process(runner)
}
