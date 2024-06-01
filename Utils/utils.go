package utils

import (
	"fmt"
	"reflect"
	"sync"
)

type Cache struct {
	cacheMap sync.Map
}

func (c *Cache) Set(key string, value any) {
	fmt.Println("SET--method")
	fmt.Println(key)
	fmt.Println(value)
	c.cacheMap.Store(key, value)
}
func (c *Cache) Get(key string) (any, bool) {
	val, ok := c.cacheMap.Load(key)
	fmt.Println("Get Method ----")
	fmt.Println(key)
	fmt.Println(val)
	if ok {
		return val, ok
	} else {
		return "", false
	}
}

func IsNil(v interface{}) bool {
	// reflect.ValueOf will panic on a nil value.
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsZero()
}
