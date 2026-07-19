package kvcache

type Cache struct {
	// NOTE@maizdrehaan: Not accessible outside the package
	cache map[string][]byte
}

func (c *Cache) Set(key string, value []byte) {

	if c.cache == nil {
		c.cache = make(map[string][]byte)
	}

	// TODO@mazidrehaan: since slices are reference types,
	// we need to make a copy of the value before storing
	// in the cache
	c.cache[key] = value
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, exists := c.cache[key]
	return value, exists
}
