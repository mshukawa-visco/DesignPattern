package singleton

import (
	"time"
)

//パッケージ外から直接呼ばない
type singleton struct {
	id int64
}

var singletonInstance *singleton

func GetSingleton(id int64) *singleton {
	if singletonInstance == nil {
		//スレッド上での振る舞いを確認するため、模擬的に生成に時間をかける
		time.Sleep(1000)
		singletonInstance = &singleton{id: id}
	}
	return singletonInstance
}

func (s *singleton) GetId() int64 {
	return s.id
}
