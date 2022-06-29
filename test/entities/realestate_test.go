package entities

import (
	"fmt"
	"realestatetrans/internal/entities"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {

	realEstate := entities.RealEstate{}
	typeOfRealEstate := reflect.TypeOf(realEstate).String()

	if typeOfRealEstate != "entities.RealEstate" {
		t.Fatal("Wrong Type.")
	}
	fmt.Printf("Type %s", typeOfRealEstate)
}

func TestNewRealEstate(t *testing.T) {

	realEstate := entities.NewRealEstate()
}
