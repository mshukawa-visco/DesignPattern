package singleton

import (
	"sync"
	"time"
)

//パッケージ外から直接呼ばない
type singleton struct {
	id int64
}

var singletonInstance *singleton
var mutex = &sync.Mutex{}

func GetSingleton(id int64) *singleton {
	// 必要以上にLockしない。
	if singletonInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singletonInstance == nil {
			//スレッド上での振る舞いを確認するため、模擬的に生成に時間をかける
			time.Sleep(1000)
			singletonInstance = &singleton{id: id}
		}
	}
	return singletonInstance
}

func (s *singleton) GetId() int64 {
	return s.id
}
