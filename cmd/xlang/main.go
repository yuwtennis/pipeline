package main

import (
	"pipelines/internal/pipelines"
	"pipelines/internal/runners"
)

func main() {
	p := pipelines.XLangElasticsearch{
		ExpansionAddr: "localhost:20000",
		EsNodeAddrs:   []string{"192.168.11.2:9200"},
		IndexName:     "MyIndex",
		MappingType:   "_doc",
	}

	var runner runners.Runner
	runner = runners.DataFlow{}

	p.Process(runner)
}
