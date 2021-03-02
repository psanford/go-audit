package main

import "github.com/golang/groupcache/lru"

type Cache interface {
	Add(lru.Key, interface{})
	Get(lru.Key) (interface{}, bool)
}

type NoCache struct{}

func (NoCache) Add(lru.Key, interface{})        {}
func (NoCache) Get(lru.Key) (interface{}, bool) { return nil, false }

// NewCache returns an lru.Cache if size is >0, NoCache otherwise
func NewCache(size int) Cache {
	if size > 0 {
		return lru.New(size)
	}
	return NoCache{}
}

func cacheSize(c Cache) int {
	switch x := c.(type) {
	case *lru.Cache:
		return x.MaxEntries
	}
	return 0
}
