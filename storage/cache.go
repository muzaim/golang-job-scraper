package storage

type Cache struct {
	data map[string]bool
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]bool),
	}
}

func (c *Cache) Exists(key string) bool {
	return c.data[key]
}

func (c *Cache) Add(key string) {
	c.data[key] = true
}
