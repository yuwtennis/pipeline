package pipelines

import (
	"github.com/stretchr/testify/assert"
	"pipelines/internal/elements"
	"pipelines/internal/pipelines"
	"testing"
)

var (
	data = "1,宅地(土地),住宅地,13115,東京都,杉並区,阿佐谷北,阿佐ケ谷,10,50000000,1900000,,85,590000,ほぼ長方形,8.2,,,,,住宅,西,私道,4,第１種低層住居専用地域,50,100,2021年第４四半期,,私道を含む取引"
)

func TestToRealEstateFn(t *testing.T) {
	result := pipelines.ToRealEstateFn(data)

	assert.IsType(t, elements.RealEstate{}, result)
}
