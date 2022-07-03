package elements

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pipelines/internal/elements"
	"reflect"
	"strconv"
	"testing"
)

func TestType(t *testing.T) {

	realEstate := elements.RealEstate{}
	typeOfRealEstate := reflect.TypeOf(realEstate).String()

	if typeOfRealEstate != "elements.RealEstate" {
		t.Fatal("Wrong Type.")
	}
	fmt.Printf("Type %s", typeOfRealEstate)
}

func TestNewRealEstate(t *testing.T) {

	var intResult int
	var float64Result float64
	valueNotEqualErr := "Two values should be the same"

	expect := map[string]string{
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

	realEstate := elements.NewRealEstate(expect)
	typeOfRealEstate := reflect.TypeOf(*realEstate).String()

	assert.Equal(t, typeOfRealEstate, "elements.RealEstate", "Wrong Type.")
	assert.Equal(t, expect["Type"], realEstate.Type, valueNotEqualErr)
	assert.Equal(t, expect["LandType"], realEstate.LandType, valueNotEqualErr)
	assert.Equal(t, expect["CityId"], realEstate.CityId, valueNotEqualErr)
	assert.Equal(t, expect["PrefectureName"], realEstate.PrefectureName, valueNotEqualErr)
	assert.Equal(t, expect["DistrictName"], realEstate.DistrictName, valueNotEqualErr)
	assert.Equal(t, expect["CityName"], realEstate.CityName, valueNotEqualErr)
	assert.Equal(t, expect["ClosestStationName"], realEstate.ClosestStationName, valueNotEqualErr)
	assert.Equal(t, expect["DurationToClosestStationInMin"], realEstate.DurationToClosestStationInMin, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["ClosedPrice"])
	assert.Equal(t, intResult, realEstate.ClosedPrice, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["UnitPriceOfFloorspace"])
	assert.Equal(t, intResult, realEstate.UnitPriceOfFloorspace, valueNotEqualErr)
	assert.Equal(t, expect["FloorPlan"], realEstate.FloorPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["AreaInSquareMeter"])
	assert.Equal(t, intResult, realEstate.AreaInSquareMeter, valueNotEqualErr)
	assert.Equal(t, expect["ShapeOfLand"], realEstate.ShapeOfLand, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(expect["FacadeInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FacadeInMeters, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["AreaRatio"])
	assert.Equal(t, intResult, realEstate.AreaRatio, valueNotEqualErr)
	assert.Equal(t, expect["YearBuilt"], realEstate.YearBuilt, valueNotEqualErr)
	assert.Equal(t, expect["ArchitectureType"], realEstate.ArchitectureType, valueNotEqualErr)
	assert.Equal(t, expect["FuturePurpose"], realEstate.FuturePurpose, valueNotEqualErr)
	assert.Equal(t, expect["FrontRoadDirection"], realEstate.FrontRoadDirection, valueNotEqualErr)
	assert.Equal(t, expect["FrontRoadType"], realEstate.FrontRoadType, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(expect["FrontRoadWithInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FrontRoadWithInMeters, valueNotEqualErr)
	assert.Equal(t, expect["CityPlan"], realEstate.CityPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["BuildingToLandRatio"])
	assert.Equal(t, intResult, realEstate.BuildingToLandRatio, valueNotEqualErr)
	intResult, _ = strconv.Atoi(expect["FloorToLandRatio"])
	assert.Equal(t, intResult, realEstate.FloorToLandRatio, valueNotEqualErr)
	assert.Equal(t, expect["AgreementDate"], realEstate.AgreementDate, valueNotEqualErr)
	assert.Equal(t, expect["RefurbishmentState"], realEstate.RefurbishmentState, valueNotEqualErr)
	assert.Equal(t, expect["AgreementNote"], realEstate.AgreementNote, valueNotEqualErr)

}
