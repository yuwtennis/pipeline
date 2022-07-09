package elements

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"pipelines/internal/elements"
	"reflect"
	"strconv"
	"testing"
)

var (
	data = map[string]string{
		"Type":                          "宅地(土地)",
		"LandType":                      "住宅地",
		"CityId":                        "13115",
		"PrefectureName":                "東京都",
		"DistrictName":                  "杉並区",
		"CityName":                      "阿佐谷北",
		"ClosestStationName":            "阿佐ケ谷",
		"DurationToClosestStationInMin": "10",
		"ClosedPrice":                   "50000000",
		"UnitPriceOfFloorspace":         "1900000",
		"FloorPlan":                     "",
		"AreaInSquareMeter":             "85",
		"UnitPriceOfSquareMeter":        "590000",
		"ShapeOfLand":                   "ほぼ長方形",
		"FacadeInMeters":                "8.2",
		"AreaRatio":                     "",
		"YearBuilt":                     "",
		"ArchitectureType":              "",
		"Purpose":                       "",
		"FuturePurpose":                 "住宅",
		"FrontRoadDirection":            "西",
		"FrontRoadType":                 "私道",
		"FrontRoadWithInMeters":         "4",
		"CityPlan":                      "第１種低層住居専用地域",
		"BuildingToLandRatio":           "50",
		"FloorToLandRatio":              "100",
		"AgreementDate":                 "2021年第４四半期",
		"RefurbishmentState":            "",
		"AgreementNote":                 "",
	}
)

func TestType(t *testing.T) {

	realEstate := elements.RealEstate{}
	typeOfRealEstate := reflect.TypeOf(realEstate).String()

	if typeOfRealEstate != "elements.RealEstate" {
		t.Fatal("Wrong Type.")
	}
	fmt.Printf("Type %s", typeOfRealEstate)
}

func TestNew(t *testing.T) {

	var intResult int
	var float64Result float64
	valueNotEqualErr := "Two values should be the same"

	realEstate := elements.New(data)
	typeOfRealEstate := reflect.TypeOf(*realEstate).String()

	assert.Equal(t, typeOfRealEstate, "elements.RealEstate", "Wrong Type.")
	assert.Equal(t, data["Type"], realEstate.Type, valueNotEqualErr)
	assert.Equal(t, data["LandType"], realEstate.LandType, valueNotEqualErr)
	assert.Equal(t, data["CityId"], realEstate.CityId, valueNotEqualErr)
	assert.Equal(t, data["PrefectureName"], realEstate.PrefectureName, valueNotEqualErr)
	assert.Equal(t, data["DistrictName"], realEstate.DistrictName, valueNotEqualErr)
	assert.Equal(t, data["CityName"], realEstate.CityName, valueNotEqualErr)
	assert.Equal(t, data["ClosestStationName"], realEstate.ClosestStationName, valueNotEqualErr)
	assert.Equal(t, data["DurationToClosestStationInMin"], realEstate.DurationToClosestStationInMin, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["ClosedPrice"])
	assert.Equal(t, intResult, realEstate.ClosedPrice, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["UnitPriceOfFloorspace"])
	assert.Equal(t, intResult, realEstate.UnitPriceOfFloorspace, valueNotEqualErr)
	assert.Equal(t, data["FloorPlan"], realEstate.FloorPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["AreaInSquareMeter"])
	assert.Equal(t, intResult, realEstate.AreaInSquareMeter, valueNotEqualErr)
	assert.Equal(t, data["ShapeOfLand"], realEstate.ShapeOfLand, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(data["FacadeInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FacadeInMeters, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["AreaRatio"])
	assert.Equal(t, intResult, realEstate.AreaRatio, valueNotEqualErr)
	assert.Equal(t, data["YearBuilt"], realEstate.YearBuilt, valueNotEqualErr)
	assert.Equal(t, data["ArchitectureType"], realEstate.ArchitectureType, valueNotEqualErr)
	assert.Equal(t, data["FuturePurpose"], realEstate.FuturePurpose, valueNotEqualErr)
	assert.Equal(t, data["FrontRoadDirection"], realEstate.FrontRoadDirection, valueNotEqualErr)
	assert.Equal(t, data["FrontRoadType"], realEstate.FrontRoadType, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(data["FrontRoadWithInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FrontRoadWithInMeters, valueNotEqualErr)
	assert.Equal(t, data["CityPlan"], realEstate.CityPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["BuildingToLandRatio"])
	assert.Equal(t, intResult, realEstate.BuildingToLandRatio, valueNotEqualErr)
	intResult, _ = strconv.Atoi(data["FloorToLandRatio"])
	assert.Equal(t, intResult, realEstate.FloorToLandRatio, valueNotEqualErr)
	assert.Equal(t, data["AgreementDate"], realEstate.AgreementDate, valueNotEqualErr)
	assert.Equal(t, data["RefurbishmentState"], realEstate.RefurbishmentState, valueNotEqualErr)
	assert.Equal(t, data["AgreementNote"], realEstate.AgreementNote, valueNotEqualErr)

}

func TestToByte(t *testing.T) {
	var result map[string]interface{}
	realEstate := elements.New(data)

	bytes := realEstate.ToByte()
	err := json.Unmarshal(bytes, &result)

	fmt.Println(*realEstate)
	fmt.Printf("%+v", result)

	assert.Nil(t, err)
	assert.Equal(t, data["Type"], result["Type"])
}
