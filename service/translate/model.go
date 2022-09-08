package translate

import (
	"time"

	"github.com/ider-zh/go-translater/baidu"
	"github.com/jellydator/ttlcache/v3"
	"gorm.io/gorm"
)

type Translate struct {
	baiduTranslate *baidu.BaiduTranslate
	database       *gorm.DB
	ttlCache       *ttlcache.Cache[string, string]
}

type TransToZhModel struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Src       string    `gorm:"primaryKey"`
	Dst       string
}
