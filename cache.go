package memoryCache

import (
	"fmt"
	"sync"
)

type Caches struct {
	mu       sync.Mutex
	innerMap map[string]int
}

func New() *Caches {
	return &Caches{innerMap: make(map[string]int)}
}

func (c *Caches) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.innerMap[key] = value
}

func (c *Caches) Get(key string) (int, error) {
	val, ok := c.innerMap[key]

	if ok {
		return val, nil
	}
	return 0, fmt.Errorf("not found key")
}

func (c *Caches) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.innerMap, key)
}
