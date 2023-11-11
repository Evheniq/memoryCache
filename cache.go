package memoryCache

import "fmt"

type Caches struct {
	innerMap map[string]int
}

func New() *Caches {
	return &Caches{innerMap: make(map[string]int)}
}

func (c *Caches) Set(key string, value int) {
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
	delete(c.innerMap, key)
}
