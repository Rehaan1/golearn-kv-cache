package kvcache

type Cache struct {
	// NOTE@mazidrehaan: Not accessible outside the package
	cache map[string][]byte
}

func (c *Cache) Set(key string, value []byte) {

	if c.cache == nil {
		c.cache = make(map[string][]byte)
	}

	// NOTE@mazidrehaan: prevents caller value
	// from being modified after Set() is called
	copied := make([]byte, len(value))
	copy(copied, value)

	c.cache[key] = copied
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, exists := c.cache[key]

	// NOTE@mazidrehaan: prevents caller from
	// modifying the internal value of the cache
	// after Get() is called
	copied := make([]byte, len(value))
	copy(copied, value)

	if !exists {
		return nil, false
	}

	return copied, exists
}
