package elements

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"pipelines/internal/elements"
	"pipelines/test"
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

func TestNew(t *testing.T) {

	var intResult int
	var float64Result float64
	valueNotEqualErr := "Two values should be the same"

	realEstate := elements.New(test.MapRealEstate)
	typeOfRealEstate := reflect.TypeOf(*realEstate).String()

	assert.Equal(t, typeOfRealEstate, "elements.RealEstate", "Wrong Type.")
	assert.Equal(t, test.MapRealEstate["Type"], realEstate.Type, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["LandType"], realEstate.LandType, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["CityId"], realEstate.CityId, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["PrefectureName"], realEstate.PrefectureName, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["DistrictName"], realEstate.DistrictName, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["CityName"], realEstate.CityName, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["ClosestStationName"], realEstate.ClosestStationName, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["DurationToClosestStationInMin"], realEstate.DurationToClosestStationInMin, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["ClosedPrice"])
	assert.Equal(t, intResult, realEstate.ClosedPrice, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["UnitPriceOfFloorspace"])
	assert.Equal(t, intResult, realEstate.UnitPriceOfFloorspace, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["FloorPlan"], realEstate.FloorPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["AreaInSquareMeter"])
	assert.Equal(t, intResult, realEstate.AreaInSquareMeter, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["ShapeOfLand"], realEstate.ShapeOfLand, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(test.MapRealEstate["FacadeInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FacadeInMeters, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["AreaRatio"])
	assert.Equal(t, intResult, realEstate.AreaRatio, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["YearBuilt"], realEstate.YearBuilt, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["ArchitectureType"], realEstate.ArchitectureType, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["FuturePurpose"], realEstate.FuturePurpose, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["FrontRoadDirection"], realEstate.FrontRoadDirection, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["FrontRoadType"], realEstate.FrontRoadType, valueNotEqualErr)
	float64Result, _ = strconv.ParseFloat(test.MapRealEstate["FrontRoadWidthInMeters"], 64)
	assert.Equal(t, float64Result, realEstate.FrontRoadWidthInMeters, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["CityPlan"], realEstate.CityPlan, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["BuildingToLandRatio"])
	assert.Equal(t, intResult, realEstate.BuildingToLandRatio, valueNotEqualErr)
	intResult, _ = strconv.Atoi(test.MapRealEstate["FloorToLandRatio"])
	assert.Equal(t, intResult, realEstate.FloorToLandRatio, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["AgreementPointOfTime"], realEstate.AgreementPointOfTime, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["RefurbishmentState"], realEstate.RefurbishmentState, valueNotEqualErr)
	assert.Equal(t, test.MapRealEstate["AgreementNote"], realEstate.AgreementNote, valueNotEqualErr)

}

func TestToByte(t *testing.T) {
	var result map[string]interface{}
	realEstate := elements.New(test.MapRealEstate)

	bytes := realEstate.ToByte()
	err := json.Unmarshal(bytes, &result)

	fmt.Println(*realEstate)
	fmt.Printf("%+v", result)

	assert.Nil(t, err)
	assert.Equal(t, test.MapRealEstate["Type"], result["Type"])
}
