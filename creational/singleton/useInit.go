package singleton

type useInit struct {
	id int64
}

var useInitInstance *useInit

// main関数が呼ばれるよりも先に呼ばれる
func init() {
	useInitInstance = &useInit{id: 0}
}

func GetUseInit() *useInit {
	return useInitInstance
}

func (ui *useInit) GetId() int64 {
	return ui.id
}

// ぶっちゃけ上のinitくらいなら次のように置き換えてもよい
var another = &useInit{id: 1}

func GetAnother() *useInit {
	return another
}
