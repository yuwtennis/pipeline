package elements

import (
	"encoding/json"
	"strconv"
)

type RealEstate struct {
	// 種類
	Type string `beam:"type"`
	// 地域
	LandType string `beam:"land_type"`
	// 市区町村コード
	CityId string `beam:"city_id"`
	// 都道府県名
	PrefectureName string `beam:"prefecture_name"`
	// 市区町村名
	DistrictName string `beam:"district_name"`
	// 地区名
	CityName string `beam:"city_name"`
	// 最寄駅：名称
	ClosestStationName string `beam:"closest_station__name"`
	// 最寄駅：距離（分）
	DurationToClosestStationInMin string `beam:"duration_to_closest_station_in_min"`
	// 取引価格（総額）
	ClosedPrice int `beam:"closed_price"`
	// 坪単価
	UnitPriceOfFloorspace int `beam:"unit_price_of_floorspace"`
	// 間取り
	FloorPlan string `beam:"floor_plan"`
	// 面積（㎡）
	AreaInSquareMeter int `beam:"area_in_square_meter"`
	// 取引価格（㎡単価）
	UnitPriceOfSquareMeter int `beam:"unit_price_of_square_meter"`
	// 土地の形状
	ShapeOfLand string `beam:"shape_of_land"`
	// 間口
	FacadeInMeters float64 `beam:"facade_in_meters"`
	// 延床面積（％）
	AreaRatio int `beam:"area_ratio"`
	// 建築年
	YearBuilt string `beam:"year_build"`
	// 築年数
	YearsSinceBuilt int `beam:"years_since_built"`
	// 建物の構造
	ArchitectureType string `beam:"architecture_type"`
	// 用途
	Purpose string `beam:"purpose"`
	// 今後の利用目的
	FuturePurpose string `beam:"future_purpose"`
	// 前面道路：方位
	FrontRoadDirection string `beam:"front_road_direction"`
	// 前面道路：種類
	FrontRoadType string `beam:"front_road_type"`
	// 前面道路：幅員（ｍ）
	FrontRoadWidthInMeters float64 `beam:"prefect"`
	// 都市計画
	CityPlan string `beam:"city_plan"`
	// 建ぺい率（％）
	BuildingToLandRatio int `beam:"building_to_land_ratio"`
	// 容積率（％）
	FloorToLandRatio int `beam:"floor_to_land_ratio"`
	// 取引時点
	AgreementPointOfTime string `beam:"agreement_point_of_time"`
	// 改装
	RefurbishmentState string `beam:"refurbishment_state"`
	//取引の事情等
	AgreementNote string `beam:"agreement_note"`
}

func New(m map[string]string) *RealEstate {
	r := new(RealEstate)

	r.Type = m["Type"]
	r.LandType = m["LandType"]
	r.CityId = m["CityId"]
	r.PrefectureName = m["PrefectureName"]
	r.DistrictName = m["DistrictName"]
	r.CityName = m["CityName"]
	r.ClosestStationName = m["ClosestStationName"]
	r.DurationToClosestStationInMin = m["DurationToClosestStationInMin"]
	r.ClosedPrice, _ = strconv.Atoi(m["ClosedPrice"])
	r.UnitPriceOfFloorspace, _ = strconv.Atoi(m["UnitPriceOfFloorspace"])
	r.FloorPlan = m["FloorPlan"]
	r.AreaInSquareMeter, _ = strconv.Atoi(m["AreaInSquareMeter"])
	r.UnitPriceOfSquareMeter, _ = strconv.Atoi(m["UnitPriceOfSquareMeter"])
	r.ShapeOfLand = m["ShapeOfLand"]
	r.FacadeInMeters, _ = strconv.ParseFloat(m["FacadeInMeters"], 64)
	r.AreaRatio, _ = strconv.Atoi(m["AreaRatio"])
	r.YearBuilt = m["YearBuilt"]
	r.ArchitectureType = m["ArchitectureType"]
	r.Purpose = m["Purpose"]
	r.FuturePurpose = m["FuturePurpose"]
	r.FrontRoadDirection = m["FrontRoadDirection"]
	r.FrontRoadType = m["FrontRoadType"]
	r.FrontRoadWidthInMeters, _ = strconv.ParseFloat(m["FrontRoadWidthInMeters"], 64)
	r.CityPlan = m["CityPlan"]
	r.BuildingToLandRatio, _ = strconv.Atoi(m["BuildingToLandRatio"])
	r.FloorToLandRatio, _ = strconv.Atoi(m["FloorToLandRatio"])
	r.AgreementPointOfTime = m["AgreementPointOfTime"]
	r.RefurbishmentState = m["RefurbishmentState"]
	r.AgreementNote = m["AgreementNote"]

	return r
}

func (re *RealEstate) ToByte() []byte {

	data, _ := json.Marshal(*re)

	return data
}
