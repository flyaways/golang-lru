golang-lru
==========

* The `lru` package which implements a fixed-size.
* Thread safe LRU cache. 
* Based on the cache in google [Groupcache lru](https://github.com/golang/groupcache/tree/master/lru).
* Simple switching between LRU,Cache,TwoQueueCache and ARCCache.

# Example

Using the LRU:

```go
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

```

* More examples can be found at [github.com/flyaways/golang-lru/_examples](https://github.com/flyaways/golang-lru/_examples).


# Reference

* [https://github.com/hashicorp/golang-lru](https://github.com/hashicorp/golang-lru)
* [https://github.com/golang/groupcache/lru](https://github.com/golang/groupcache/tree/master/lru)

# Copyright and License
Copyright 2018 The golang-lru Authors. All rights reserved.

for the golang-lru Authors. Code is released under
[the Apache 2 license](https://github.com/flyaways/golang-lru/blob/master/LICENSE).
