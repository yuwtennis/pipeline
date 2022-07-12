package test

var (
	MapRealEstate = map[string]string{
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
		"YearBuilt":                     "昭和56年",
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
	CSVRealEstate = []string{
		"1,宅地(土地),住宅地,13115,東京都,杉並区,阿佐谷北,阿佐ケ谷,10,50000000,1900000,,85,590000,ほぼ長方形,8.2,,昭和56年,,,住宅,西,私道,4,第１種低層住居専用地域,50,100,2021年第４四半期,,私道を含む取引",
	}
)
