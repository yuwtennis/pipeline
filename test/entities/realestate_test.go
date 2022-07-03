package entities

import (
	"fmt"
	"log"
	"pipelines/internal/elements"
	"reflect"
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

	data := map[string]string{
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

	realEstate := elements.NewRealEstate(data)
	log.Println(*realEstate)
	typeOfRealEstate := reflect.TypeOf(*realEstate).String()

	if typeOfRealEstate != "elements.RealEstate" {
		t.Fatal("Wrong Type.")
	}

}
