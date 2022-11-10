package pipelines

import "pipelines/internal/runners"

type Pipeline interface {
	Process(runner runners.Runner)
}
