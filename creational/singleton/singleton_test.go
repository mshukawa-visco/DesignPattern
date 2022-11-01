package singleton

import (
	"testing"
)

func TestGetSingleton(t *testing.T) {
	instance1 := GetSingleton(3)
	if instance1 == nil {
		t.Fatal("GetSingleton(): expected pointer is not nil")
	}

	if id := instance1.GetId(); id != 3 {
		t.Errorf("GetId(): expected=%d, but got=%d", 3, id)
	}

	// インスタンス再取得
	instance2 := GetSingleton(0)
	if instance2 != instance1 {
		t.Error("expected same instance")
	}

	if id := instance2.GetId(); id != 3 {
		t.Errorf("GetId(): expected=%d, but got=%d", 3, id)
	}
}
