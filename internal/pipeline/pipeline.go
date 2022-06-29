package pipeline

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	csvmap "github.com/recursionpharma/go-csv-map"
	"realestatetrans/internal"
	"realestatetrans/internal/entities"
	"regexp"
	"strings"
)

func strToDateFn(event *entities.RealEstate) *entities.RealEstate {
	regex := regexp.MustCompile(`(\d+)第(.)半期`)
	res := regex.FindAllStringSubmatch(event.AgreementDate, -1)
	year := res[0][1]
	quarter := res[0][2]

	switch quarter {
	case "１":
		event.AgreementDate = fmt.Sprintf("%s-%s-1", year, "4")
	case "２":
		event.AgreementDate = fmt.Sprintf("%s-%s-1", year, "7")
	case "３":
		event.AgreementDate = fmt.Sprintf("%s-%s-1", year, "10")
	case "４":
		event.AgreementDate = fmt.Sprintf("%s-%s-1", year, "1")
	}

	return event
}

func toRealEstateFn(event string) *entities.RealEstate {
	keys := []string{"Type",
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

	return entities.NewRealEstate(record[0])
}

func Run(events []string) {
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()
	beam.Create(s, events)

	// TODO To RealEstate
	// TODO 建築年 & 取引時点 to time object

	err := beamx.Run(context.Background(), p)

	internal.Check(err)
}
