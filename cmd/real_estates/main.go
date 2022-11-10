package main

import (
	"pipelines/internal/pipelines"
	"pipelines/internal/runners"
)

func main() {
	p := pipelines.RealEstate{}
	var runner runners.Runner
	runner = runners.Direct{}

	p.Process(runner)
}
