package singleton

// package外からはstructを呼ばない
type singletonCounter struct {
	count int
}

var counter *singletonCounter

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
