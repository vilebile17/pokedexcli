package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.locations[key]
	if ok {
		stuff := entry.val
		return stuff, true
	}
	return nil, false
}
