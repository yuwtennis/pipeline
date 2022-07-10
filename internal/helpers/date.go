package helpers

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

const (
	ShowEraName          = "昭和"
	HeiseiEraName        = "平成"
	ReiwaEraName         = "令和"
	BeforeWarTimeEraName = "戦前"
	FirstQuarterName     = "第１四半期"
	SecondQuarterName    = "第２四半期"
	ThirdQuarterName     = "第３四半期"
	FourthQuarterName    = "第４四半期"
)

var (
	regexpEra     = regexp.MustCompile(`^(?P<Era>.*?)(?P<Year>\d+)?年?$`)
	regexpQuarter = regexp.MustCompile(`^(?P<Year>\d+)年(?P<Quarter>.*)$`)
)

//JpnEraToISO8601Str Converts japanese regnal era such as "昭和２年" to ISO8601 format string
func JpnEraToISO8601Str(s string) string {
	var (
		year    int
		result  []string
		eraName string
		eraYear int
	)

	result = regexpEra.FindAllStringSubmatch(s, -1)[0]
	eraName = result[1]

	if s != BeforeWarTimeEraName {
		eraYear, _ = strconv.Atoi(result[2])
	}

	switch eraName {
	case ShowEraName:
		year = 1926 + eraYear - 1
	case HeiseiEraName:
		year = 1989 + eraYear - 1
	case ReiwaEraName:
		year = 2019 + eraYear - 1
	case BeforeWarTimeEraName:
		year = 1945
	default:
		log.Panicln("Invalid era.")
	}

	t := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	return t.Format(time.RFC3339)
}

//JpnQuarterToISO8601Str Converts Japanese quarter such as "2021年第１四半期" to ISO8601 format string
func JpnQuarterToISO8601Str(s string) string {
	var (
		t     time.Time
		month time.Month
	)
	result := regexpQuarter.FindAllStringSubmatch(s, -1)[0]
	year, _ := strconv.Atoi(result[1])
	quarterName := result[2]

	switch quarterName {
	case FirstQuarterName:
		month = time.April
	case SecondQuarterName:
		month = time.July
	case ThirdQuarterName:
		month = time.October
	case FourthQuarterName:
		year = year + 1
		month = time.January
	default:
		log.Panicln("Invalid quarter.")
	}

	t = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	return t.Format(time.RFC3339)
}
