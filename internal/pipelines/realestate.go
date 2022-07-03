package pipelines

import (
	"bytes"
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	csvmap "github.com/recursionpharma/go-csv-map"
	"pipelines/internal"
	"pipelines/internal/elements"
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
	internal.Check(err)

	return elements.NewRealEstate(record[0])
}

func (re *RealEstate) Process(elements []string) {
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()
	beam.Create(s, elements)

	// TODO To RealEstate
	// TODO 建築年 & 取引時点 to time object

	err := beamx.Run(context.Background(), p)

	internal.Check(err)
}
