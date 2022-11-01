package singleton

import (
	"sync"
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

func TestMultiThread(t *testing.T) {
	var ids [100]int64

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int64) {
			defer wg.Done()
			ids[id] = GetSingleton(id).GetId()
		}(&wg, int64(i))
	}
	wg.Wait()

	for i := 1; i < 100; i++ {
		if ids[0] != ids[i] {
			t.Errorf("Got Multiple Ids. ids[0]=%d, ids[%d]=%d", ids[0], i, ids[i])
		}
	}
}
