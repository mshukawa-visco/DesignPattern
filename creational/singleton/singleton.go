package singleton

//パッケージ外から直接呼ばない
type singleton struct {
	id int64
}

var singletonInstance *singleton

func GetSingleton(id int64) *singleton {
	if singletonInstance == nil {
		singletonInstance = &singleton{id: id}
	}
	return singletonInstance
}

func (s *singleton) GetId() int64 {
	return s.id
}
