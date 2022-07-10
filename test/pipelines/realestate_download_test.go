package pipelines

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
	"pipelines/internal/elements"
	"pipelines/internal/pipelines"
	"pipelines/test"
	"testing"
)

func TestToRealEstateFn(t *testing.T) {
	expect := elements.New(test.MapRealEstate)

	pipeline, scope, collection := ptest.CreateList(test.CSVRealEstate)
	result := beam.ParDo(scope, pipelines.ToRealEstateFn, collection)

	ptest.Run(pipeline)
	passert.Equals(scope, result, expect)
}

func TestToTimestampFn(t *testing.T) {
	expect := elements.New(test.MapRealEstate)
	expect.YearBuilt = "1981-01-01T00:00:00Z"
	expect.AgreementPointOfTime = "2021-04-01T00:00:00Z"

	data := []*elements.RealEstate{
		elements.New(test.MapRealEstate),
	}
	pipeline, scope, collection := ptest.CreateList(data)
	result := beam.ParDo(scope, pipelines.ToTimestampFn, collection)

	ptest.Run(pipeline)
	passert.Equals(scope, result, expect)
}
