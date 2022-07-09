package helpers

import (
	"regexp"
	"strconv"
	"time"
)

var (
	eraMap = map[string]int{
		"昭和": 1926,
		"平成": 1989,
		"令和": 2019,
	}
	quarterMap = map[string]time.Month{
		"第１四半期": time.April,
		"第２四半期": time.July,
		"第３四半期": time.October,
		"第４四半期": time.January,
	}

	regexpEra     = regexp.MustCompile(`^(?P<Era>.*?)(?P<Year>\d+).*$`)
	regexpQuarter = regexp.MustCompile(`^(?P<Year>\d+)年(?P<Quarter>.*)$`)
)

func JpnEraToISO8601Str(s string) string {
	result := regexpEra.FindAllStringSubmatch(s, -1)[0]

	era := eraMap[result[1]]
	year, _ := strconv.Atoi(result[2])

	t := time.Date(era+year, time.January, 1, 0, 0, 0, 0, time.UTC)

	return t.Format(time.RFC3339)
}

func JpnQuarterToISO8601Str(s string) string {
	var (
		t     time.Time
		month time.Month
	)
	result := regexpQuarter.FindAllStringSubmatch(s, -1)[0]
	year, _ := strconv.Atoi(result[1])

	switch result[2] {
	case "第１四半期":
		month = time.April
	case "第２四半期":
		month = time.July
	case "第３四半期":
		month = time.October
	case "第４四半期":
		year = year + 1
		month = time.January
	}

	t = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	return t.Format(time.RFC3339)
}
