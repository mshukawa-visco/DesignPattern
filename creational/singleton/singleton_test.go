package singleton

import (
	"testing"
)

//Singletonは他クラスで使う場合も含め単体テストし辛い点に留意
func TestGetCounter(t *testing.T) {
	// インスタンス取得
	counter1 := GetCounter()
	if counter1 == nil {
		t.Fatal("GetCounter(): expected pointer is not nil")
	}

	expectedCounter := counter1
	if current := counter1.Increment(); current != 1 {
		t.Errorf("Increment(): expected=%d, but got=%d", 1, current)
	}

	// インスタンス再取得
	counter2 := GetCounter()
	if counter2 != expectedCounter {
		t.Error("expected same instance")
	}

	if current := counter2.Decrement(); current != 0 {
		t.Errorf("Decrement(): expected=%d, but got=%d", 0, current)
	}
}
