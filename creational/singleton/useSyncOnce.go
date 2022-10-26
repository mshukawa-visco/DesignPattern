package singleton

import (
	"sync"
)

type singletonDB struct {
	name string
}

var once sync.Once
var database *singletonDB

func GetDataBase(dbName string) *singletonDB {
	once.Do(func() {
		database = &singletonDB{name: dbName}
	})
	return database
}

func (db *singletonDB) GetDBInfo() string {
	return db.name
}
