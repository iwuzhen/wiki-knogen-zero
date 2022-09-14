package freshstorage

import (
	"errors"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	queryLock sync.Mutex
)

func NewFreshStorage(postgresConn string) *FreshStorage {

	storageDB, err := gorm.Open(postgres.Open(postgresConn), &gorm.Config{})
	if err != nil {
		log.Fatal("postgres cnnect fail", err)
	}
	storageDB.AutoMigrate(&StorageBucket{})
	// cace one minute

	return &FreshStorage{storageDB}

}

func (c *FreshStorage) GetLastestRecord(path, key string) *StorageBucket {
	var sb StorageBucket
	c.database.Where("path = ? AND key = ?", path, key).Last(&sb)
	return &sb
}

func (c *FreshStorage) PutRecord(path, key, data string) *StorageBucket {
	// dataJSon, err := json.Marshal(data)
	// if err != nil {
	// 	log.Println("json Marshal error", err)
	// }
	dataitem := StorageBucket{
		Path: path,
		Key:  key,
		// Data: datatypes.JSON(dataJSon),
		Data: data,
	}
	c.database.Create(&dataitem)
	return &dataitem
}

func (c *FreshStorage) PutConsecutiveRecord(path, key string, ID uint, data string) (dataitem *StorageBucket, err error) {
	// dataJSon, err := json.Marshal(data)
	// if err != nil {
	// 	log.Println("json Marshal error", err)
	// }
	dataitem = &StorageBucket{
		Path: path,
		Key:  key,
		// Data: datatypes.JSON(dataJSon),
		Data: data,
	}

	var sb StorageBucket
	c.database.Where("path = ? AND key = ?", path, key).Last(&sb)
	if ID == sb.ID {
		c.database.Create(&dataitem)
	} else {
		err = errors.New("ID expired")
	}

	return
}
