package freshstorage

import (
	"time"

	"gorm.io/gorm"
)

type FreshStorage struct {
	database *gorm.DB
	// ttlCache *ttlcache.Cache[string, interface{}]
}

type StorageBucket struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	// UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Path      string         `gorm:"index:idx_member"` //数据类别Path
	Key       string         `gorm:"index:idx_member"` //数据类别Key
	// Data      datatypes.JSON
	Data string
}
