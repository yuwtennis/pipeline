package entities

import (
	"github.com/mitchellh/mapstructure"
	"realestatetrans/internal"
)

var HeaderToKeys = map[string]string{
	"種類":         "Type",
	"地域":         "LandType",
	"市区町村コード":    "CityId",
	"都道府県名":      "PrefectureName",
	"市区町村名":      "DistrictName",
	"地区名":        "CityName",
	"最寄駅：名称":     "ClosestStationName",
	"最寄駅：距離（分）":  "DurationToClosestStationInMin",
	"取引価格（総額）":   "ClosedPrice",
	"坪単価":        "UnitPriceOfFloorspace",
	"間取り":        "FloorPlan",
	"面積（㎡）":      "AreaInSquareMeter",
	"取引価格（㎡単価）":  "UnitPriceOfSquareMeter",
	"土地の形状":      "ShapeOfLand",
	"間口":         "FacadeInMeters",
	"延床面積（％）":    "AreaRatio",
	"建築年":        "YearBuilt",
	"建物の構造":      "ArchitectureType",
	"用途":         "Purpose",
	"今後の利用目的":    "FuturePurpose",
	"前面道路：方位":    "FrontRoadDirection",
	"前面道路：種類":    "FrontRoadType",
	"前面道路：幅員（ｍ）": "FrontRoadWithInMeters",
	"都市計画":       "CityPlan",
	"建ぺい率（％）":    "BuildingToLandRatio",
	"容積率（％）":     "FloorToLandRatio",
	"取引時点":       "AgreementDate",
	"改装":         "RefurbishmentState",
	"取引の事情等":     "AgreementNote",
}

type RealEstate struct {
	// Id
	Id string
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
	m2 := make(map[string]interface{})
	r := new(RealEstate)

	for k, v := range m {
		m2[HeaderToKeys[k]] = v
	}

	err := mapstructure.Decode(m2, &r)
	internal.Check(err)

	return r
}
