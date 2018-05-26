package main

import (
	"fmt"

	"github.com/flyaways/golang-lru"
	"github.com/flyaways/golang-lru/simplelru"
)

var (
	cache simplelru.LRUCache
	err   error
)

func main() {
	cache, err = lru.New(8)
	/*
		cache, err = simplelru.NewLRU(8, func(key interface{}, value interface{}) {
			fmt.Println(time.Now().Format(time.RFC3339Nano))
		})
		cache, err = lru.NewWithEvict(8, func(key interface{}, value interface{}) {
			fmt.Println(time.Now().Format(time.RFC3339Nano))
		})
		cache, err = lru.New2Q(8)
		cache, err = lru.New2QParams(8, 0.5, 0.5)
		cache, err = lru.NewARC(8)
	*/

	if err != nil {
		panic(err)
	}

	for i := 0; i < 16; i++ {
		fmt.Println(i, cache.Add(i, nil))
	}

	fmt.Println(cache.Get(6))
	fmt.Println(cache.Contains(6))
	fmt.Println(cache.Peek(6))
	fmt.Println(cache.GetOldest())
	cache.Remove(6)
	fmt.Println(cache.RemoveOldest())
	fmt.Println(cache.Keys())
	fmt.Println(cache.Len())
	cache.Purge()
	fmt.Println(cache.Len())
}
