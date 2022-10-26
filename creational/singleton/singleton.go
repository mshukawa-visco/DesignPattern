package singleton

// package外からはstructを呼ばない
type singletonCounter struct {
	count int
}

// インスタンスを保持する変数も小文字にすることでエクスポートを行わない
var counter *singletonCounter

// ここでインスタンスが単一であることを保証
func GetCounter() *singletonCounter {
	if counter == nil {
		counter = &singletonCounter{count: 0}
	}
	return counter
}

func (c *singletonCounter) Increment() int {
	c.count++
	return c.count
}

func (c *singletonCounter) Decrement() int {
	c.count--
	return c.count
}
