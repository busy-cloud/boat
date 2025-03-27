package lib

import (
	"sync"
	"time"
)

type cacheLoaderItem[T any] struct {
	value *T
	err   error

	expireAt int64
}

type CacheLoaderFunc[T any] func(key string) (*T, error)

type CacheLoader[T any] struct {
	items map[string]*cacheLoaderItem[T]
	lock  sync.RWMutex

	Timeout int64
	Loader  CacheLoaderFunc[T]
}

func (c *CacheLoader[T]) Invalid(key string) {
	if item, ok := c.items[key]; ok {
		item.expireAt = time.Now().Unix()
	}
}

func (c *CacheLoader[T]) Load(key string) (*T, error) {
	c.lock.RLock()

	if item, ok := c.items[key]; ok {
		if time.Now().Unix() < item.expireAt {
			c.lock.RUnlock()
			return item.value, item.err
		}
	}

	//转成写锁
	c.lock.RUnlock()
	c.lock.Lock()
	defer c.lock.Unlock()

	item := &cacheLoaderItem[T]{}

	//正式加载
	item.value, item.err = c.Loader(key)
	item.expireAt = time.Now().Unix() + c.Timeout

	return item.value, item.err
}
