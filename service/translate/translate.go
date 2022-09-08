package translate

import (
	"log"
	"time"

	"github.com/ider-zh/go-translater/baidu"
	"github.com/jellydator/ttlcache/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewTranslate(appid, secret, postgresConn string) *Translate {
	trans := baidu.NewBaiduTranslater(appid, secret, baidu.Senior, baidu.ZH)

	translateDB, err := gorm.Open(postgres.Open(postgresConn), &gorm.Config{})
	if err != nil {
		log.Fatal("postgres cnnect fail", err)
	}
	// create table
	translateDB.AutoMigrate(&TransToZhModel{})

	// 缓存一个月
	ttlCache := ttlcache.New[string, string](ttlcache.WithTTL[string, string](60 * 24 * time.Hour))

	// 2 day delete over time value
	go func() {
		for {
			time.Sleep(48 * time.Hour)
			ttlCache.DeleteExpired()
		}
	}()

	return &Translate{trans, translateDB, ttlCache}
}

func (c *Translate) TranslateSliceString(sourceString []string) []string {
	dstString := make([]string, len(sourceString))
	leakIndexSlice := []int{}
	leakSrcSlice := []string{}

	// query from cache
	for i := range sourceString {
		item := c.ttlCache.Get(sourceString[i])
		if item != nil {
			dstString[i] = item.Value()
		} else {
			leakIndexSlice = append(leakIndexSlice, i)
		}
	}

	// query from database
	for _, i := range leakIndexSlice {
		var dbRet TransToZhModel
		// result := c.database.Where(&TransToZhModel{Src: sourceString[i]}).Take(&dbRet)
		//only update one year age
		t := time.Now()
		t = t.Add(time.Hour * 24 * 30 * 12 * -1)

		result := c.database.Where("src = ? AND updated_at >= ?", sourceString[i], t).Take(&dbRet)
		if result.RowsAffected >= 1 {
			c.ttlCache.Set(sourceString[i], dbRet.Dst, ttlcache.DefaultTTL)
		} else {
			leakSrcSlice = append(leakSrcSlice, sourceString[i])
		}
	}

	// call baidu vip translate api
	if len(leakSrcSlice) > 0 {
		ret := c.baiduTranslate.Translate(leakSrcSlice)
		for i := range ret {
			c.ttlCache.Set(leakSrcSlice[i], ret[i], ttlcache.DefaultTTL)
			NewObj := TransToZhModel{
				Src: leakSrcSlice[i],
				Dst: ret[i],
			}
			c.database.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "src"}},
				DoUpdates: clause.AssignmentColumns([]string{"dst"}),
			}).Create(&NewObj)
		}
	}

	// fix leak
	for _, i := range leakIndexSlice {
		item := c.ttlCache.Get(sourceString[i])
		if item != nil {
			dstString[i] = item.Value()
		} else {
			dstString[i] = ""
		}
	}
	return dstString
}
