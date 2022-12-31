package asyncDispatching

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"PP/worker/genericMath"
)

type Cache struct {
	sync.RWMutex
	items map[string]interface{}
}

type Item struct {
	data interface{}
}

func NewCache() *Cache {
	cache := Cache{
		items: make(map[string]interface{}),
	}

	return &cache
}

func (c *Cache) SetItem(key string, item interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = item
}

func (c *Cache) GetItem(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.items) == 0 {
		return nil, ErrEmptyCache
	}

	item, found := c.items[key]

	if !found {
		return nil, ErrNoItem
	}

	return item, nil
}

func (c *Cache) GetAndDeleteItem(key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()

	if len(c.items) == 0 {
		return nil, ErrEmptyCache
	}

	item, found := c.items[key]

	if !found {
		return nil, ErrNoItem
	}

	delete(c.items, key)

	return item, nil
}

func (c *Cache) Drop() {
	c.Lock()
	defer c.Unlock()

	c.items = make(map[string]interface{})
}

func BinaryIntFuncHash(f BinaryFloatFunction, d1 float64, d2 float64) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() +
		fmt.Sprint(d1) + fmt.Sprint(d2)
}

func UnaryIntSeqFuncHash(f UnaryFloatFunction, seq *genericMath.FloatSequence) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() +
		fmt.Sprint(unsafe.Pointer(seq))
}
