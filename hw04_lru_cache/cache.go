package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var flag bool
	if c.items[key] != nil {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
		flag = true
	} else {
		item := c.queue.PushFront(value)
		c.items[key] = item
		if c.queue.Len() > c.capacity {
			for k, v := range c.items {
				if v == c.queue.Back() {
					delete(c.items, k)
					break
				}
			}
			c.queue.Remove(c.queue.Back())
		}
	}
	return flag
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if c.items[key] != nil {
		c.queue.MoveToFront(c.items[key])
		return c.items[key].Value, true
	} else {
		return nil, false
	}
}
func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
