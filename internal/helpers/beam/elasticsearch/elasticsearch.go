package elasticsearchio

const (
	writeURN            = "beam:transform:org.apache.beam:elasticsearch_write:v1"
	serviceGradleTarget = ":sdks:java:io:expansion-service:runExpansionService"
)

func Write() {

}

type writeOption func(*writePayload)

type writePayload struct {
}
