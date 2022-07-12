package pipelines

type Pipeline interface {
	Process(elements string)
}
