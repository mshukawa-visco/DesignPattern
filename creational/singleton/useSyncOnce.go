package singleton

import (
	"sync"
)

type useSyncOnce struct {
	id int64
}

var useSyncOnceInstance *useSyncOnce
var once sync.Once

func GetUseSyncOnce(id int64) *useSyncOnce {
	if useSyncOnceInstance == nil {
		once.Do(func() {
			//スレッド上での振る舞いを確認するため、模擬的に生成に時間をかける
			useSyncOnceInstance = &useSyncOnce{id: id}
		})
	}
	return useSyncOnceInstance
}

func (uso *useSyncOnce) GetId() int64 {
	return uso.id
}
