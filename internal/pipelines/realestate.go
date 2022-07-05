package pipelines

import (
	"bytes"
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	csvmap "github.com/recursionpharma/go-csv-map"
	"pipelines/internal/elements"
	"pipelines/internal/helpers"
	"pipelines/internal/runners/direct"
	"strings"
)

type RealEstate struct{}

func (re *RealEstate) toRealEstateFn(event string) *elements.RealEstate {
	keys := []string{
		"Type",
		"LandType",
		"CityId",
		"PrefectureName",
		"DistrictName",
		"CityName",
		"ClosestStationName",
		"DurationToClosestStationInMin",
		"ClosedPrice",
		"UnitPriceOfFloorspace",
		"FloorPlan",
		"AreaInSquareMeter",
		"UnitPriceOfSquareMeter",
		"ShapeOfLand",
		"FacadeInMeters",
		"AreaRatio",
		"YearBuilt",
		"ArchitectureType",
		"Purpose",
		"FuturePurpose",
		"FrontRoadDirection",
		"FrontRoadType",
		"FrontRoadWithInMeters",
		"CityPlan",
		"BuildingToLandRatio",
		"FloorToLandRatio",
		"AgreementDate",
		"RefurbishmentState",
		"AgreementNote",
	}

	buf := bytes.NewBufferString(strings.Join(keys, ",") + "\n" + event)
	reader := csvmap.NewReader(buf)
	record, err := reader.ReadAll()
	helpers.Check(err)

	return elements.NewRealEstate(record[0])
}

func (re *RealEstate) toTimestamp(e *elements.RealEstate) *elements.RealEstate {
	e.YearBuilt = "NEW_DATE"
	e.AgreementDate = "NEW_DATE"

	return e
}

func (re *RealEstate) Process(elements []string) {
	d := direct.Direct{}

	p, s := d.Init()
	lines := beam.Create(s, elements)
	entities := beam.ParDo(s, re.toRealEstateFn, lines)
	entities = beam.ParDo(s, re.toTimestamp, entities)

	err := beamx.Run(context.Background(), p)

	helpers.Check(err)
}
