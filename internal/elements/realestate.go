package elements

import (
	"strconv"
)

type RealEstate struct {
	// 種類
	Type string
	// 地域
	LandType string
	// 市区町村コード
	CityId string
	// 都道府県名
	PrefectureName string
	// 市区町村名
	DistrictName string
	// 地区名
	CityName string
	// 最寄駅：名称
	ClosestStationName string
	// 最寄駅：距離（分）
	DurationToClosestStationInMin string
	// 取引価格（総額）
	ClosedPrice int
	// 坪単価
	UnitPriceOfFloorspace int
	// 間取り
	FloorPlan string
	// 面積（㎡）
	AreaInSquareMeter int
	// 取引価格（㎡単価）
	UnitPriceOfSquareMeter int
	// 土地の形状
	ShapeOfLand string
	// 間口
	FacadeInMeters int
	// 延床面積（％）
	AreaRatio int
	// 建築年
	YearBuilt string
	// 建物の構造
	ArchitectureType string
	// 用途
	Purpose string
	// 今後の利用目的
	FuturePurpose string
	// 前面道路：方位
	FrontRoadDirection string
	// 前面道路：種類
	FrontRoadType string
	// 前面道路：幅員（ｍ）
	FrontRoadWithInMeters float64
	// 都市計画
	CityPlan string
	// 建ぺい率（％）
	BuildingToLandRatio int
	// 容積率（％）
	FloorToLandRatio int
	// 取引時点
	AgreementDate string
	// 改装
	RefurbishmentState string
	//取引の事情等
	AgreementNote string
}

func NewRealEstate(m map[string]string) *RealEstate {
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
	r.FacadeInMeters, _ = strconv.Atoi(m["FacadeInMeters"])
	r.AreaRatio, _ = strconv.Atoi(m["AreaRatio"])
	r.YearBuilt = m["YearBuilt"]
	r.ArchitectureType = m["ArchitectureType"]
	r.Purpose = m["Purpose"]
	r.FuturePurpose = m["FuturePurpose"]
	r.FrontRoadDirection = m["FrontRoadDirection"]
	r.FrontRoadType = m["FrontRoadType"]
	r.FrontRoadWithInMeters, _ = strconv.ParseFloat(m["FrontRoadWithInMeters"], 64)
	r.CityPlan = m["CityPlan"]
	r.BuildingToLandRatio, _ = strconv.Atoi(m["BuildingToLandRatio"])
	r.FloorToLandRatio, _ = strconv.Atoi(m["FloorToLandRatio"])
	r.AgreementDate = m["AgreementDate"]
	r.RefurbishmentState = m["RefurbishmentState"]
	r.AgreementNote = m["AgreementNote"]

	return r
}
