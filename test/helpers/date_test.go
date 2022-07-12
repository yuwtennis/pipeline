package helpers

import (
	"github.com/stretchr/testify/assert"
	"pipelines/internal/helpers"
	"testing"
)

const (
	strShowaEra      = "昭和56年"
	strHeiseiEra     = "平成10年"
	strReiwaEra      = "令和2年"
	strFirstQuarter  = "2021年第１四半期"
	strSecondQuarter = "2021年第２四半期"
	strThirdQuarter  = "2021年第３四半期"
	strFourthQuarter = "2021年第４四半期"
)

func TestJpnEraToISO8601Str(t *testing.T) {
	expect1 := "1981-01-01T00:00:00Z"
	result1 := helpers.JpnEraToISO8601Str(strShowaEra)

	expect2 := "1998-01-01T00:00:00Z"
	result2 := helpers.JpnEraToISO8601Str(strHeiseiEra)

	expect3 := "2020-01-01T00:00:00Z"
	result3 := helpers.JpnEraToISO8601Str(strReiwaEra)

	assert.Equal(t, expect1, result1)
	assert.Equal(t, expect2, result2)
	assert.Equal(t, expect3, result3)
}

func TestJpnQuarterToISO8601Str(t *testing.T) {
	expect1 := "2021-04-01T00:00:00Z"
	result1 := helpers.JpnQuarterToISO8601Str(strFirstQuarter)

	expect2 := "2021-07-01T00:00:00Z"
	result2 := helpers.JpnQuarterToISO8601Str(strSecondQuarter)

	expect3 := "2021-10-01T00:00:00Z"
	result3 := helpers.JpnQuarterToISO8601Str(strThirdQuarter)

	expect4 := "2022-01-01T00:00:00Z"
	result4 := helpers.JpnQuarterToISO8601Str(strFourthQuarter)

	assert.Equal(t, expect1, result1)
	assert.Equal(t, expect2, result2)
	assert.Equal(t, expect3, result3)
	assert.Equal(t, expect4, result4)

}
