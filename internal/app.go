package internal

import "pipelines/internal/pipelines"

// Run is the orchestrator of this application
func Run() {
	p := pipelines.RealEstate{}
	p.Process()
}
