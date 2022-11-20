package pipelines

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/recursionpharma/go-csv-map"
	"github.com/relvacode/iso8601"
	"os"
	"pipelines/internal/clients"
	"pipelines/internal/elements"
	"pipelines/internal/helpers"
	beam2 "pipelines/internal/helpers/beam"
	"pipelines/internal/runners"
	"strings"
	"time"
)

const (
	PREFECTURE_CODE_TOKYO = "13"
	CITY_CODE_ALL_CITY    = "13999"
)

var (
	Keys = []string{
		"Id",
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
		"FrontRoadWidthInMeters",
		"CityPlan",
		"BuildingToLandRatio",
		"FloorToLandRatio",
		"AgreementPointOfTime",
		"RefurbishmentState",
		"AgreementNote",
	}
)

type RealEstate struct{}

func init() {
	beam.RegisterFunction(ToTimestampFn)
	beam.RegisterFunction(ToRealEstateFn)
	beam.RegisterFunction(CalcYearsSinceBuilt)
	beam.RegisterFunction(beam2.ToElasticsearchFn)
}

func ToRealEstateFn(event string) *elements.RealEstate {

	event = strings.ReplaceAll(event, "\"", "")

	reader := csvmap.NewReader(bytes.NewBufferString(event))
	reader.Columns = Keys
	record, _ := reader.ReadAll()

	return elements.New(record[0])
}

func ToTimestampFn(e *elements.RealEstate) *elements.RealEstate {
	if e.YearBuilt != "" {
		e.YearBuilt = helpers.JpnEraToISO8601Str(e.YearBuilt)
	}

	if e.AgreementPointOfTime != "" {
		e.AgreementPointOfTime = helpers.JpnQuarterToISO8601Str(e.AgreementPointOfTime)
	}

	return e
}

// TODO unittest
func CalcYearsSinceBuilt(e *elements.RealEstate) *elements.RealEstate {
	now := time.Now()
	then, _ := iso8601.ParseString(e.YearBuilt)

	// TODO Calculate more specifically
	e.YearsSinceBuilt = now.Year() - then.Year()

	return e
}

func (re RealEstate) Process(runner runners.Runner) {
	var decodedSlice []string

	downloadedFile := "/tmp/file.zip"
	from := "20014"
	to := "20214"
	prefecture := PREFECTURE_CODE_TOKYO
	city := CITY_CODE_ALL_CITY

	requestUrl := fmt.Sprintf(
		"https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?DLF=true&TTC-From=%s&TTC-To=%s&TDK=%s&SKC=%s",
		from,
		to,
		prefecture,
		city)

	requestHeader := make(map[string]string)
	requestHeader["referer"] = fmt.Sprintf(
		"https://www.land.mlit.go.jp/webland/servlet/DownloadServlet?TDK=%s&SKC=%s&TDIDFrom=%s&TDIDTo=%s",
		prefecture,
		city,
		from,
		to)

	// TODO Use channel
	contents := clients.Download(requestUrl, requestHeader)

	os.WriteFile(downloadedFile, contents, 0644)
	read, err := zip.OpenReader(downloadedFile)
	helpers.Check(err)

	for _, file := range read.File {
		rawFile, err := file.Open()
		helpers.Check(err)
		scanner := bufio.NewScanner(rawFile)

		for scanner.Scan() {
			decodedString := string(helpers.ShiftJisToUTF8([]byte(scanner.Text())))
			decodedSlice = append(decodedSlice, decodedString)
		}
	}

	err = os.Remove(downloadedFile)
	helpers.Check(err)

	// First Line is a header
	decodedSlice = decodedSlice[1:]

	fmt.Printf("len: [%d], cap: [%d] \n", len(decodedSlice), cap(decodedSlice))
	pipeline, scope := runners.Init()

	lines := beam.CreateList(scope, decodedSlice)
	entities := beam.ParDo(scope, ToRealEstateFn, lines)
	entities = beam.ParDo(scope, ToTimestampFn, entities)
	entities = beam.ParDo(scope, CalcYearsSinceBuilt, entities)
	beam.ParDo(scope, beam2.ToElasticsearchFn, entities)

	_, err = runner.Execute(context.Background(), pipeline)
	helpers.Check(err)
}
