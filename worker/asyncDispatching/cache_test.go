package asyncDispatching

import (
	"testing"
)

const TestingItem1 = 1
const TestingItem2 = 2

func initTestingCache() *Cache {
	return &Cache{
		items: map[string]interface{}{
			"key1": 1,
		},
	}
}

func TestCache_GetItem(t *testing.T) {
	cache := initTestingCache()

	item, err := cache.GetItem("key1")

	if err != nil {
		t.Errorf("Got %e", err)
	} else if item != TestingItem1 {
		t.Errorf("Expected %d, got %d", TestingItem1, item)
	}

	item, err = cache.GetItem("fjfjf")

	if err != ErrNoItem {
		t.Errorf("Expected ErrNoItem, got %e", err)
	}

	_, err = (&Cache{}).GetItem("fjfjf")

	if err != ErrEmptyCache {
		t.Errorf("Expected ErrEmptyCache, got %e", err)
	}
}

func TestCache_GetAndDeleteItem(t *testing.T) {
	cache := initTestingCache()

	item, err := cache.GetAndDeleteItem("key1")

	if err != nil {
		t.Errorf("Got %e", err)
	} else if item != TestingItem1 {
		t.Errorf("Expected %d, got %d", TestingItem1, item)
	}

	_, found := cache.items["key1"]

	if found {
		t.Errorf("Expected %d, insted not found", TestingItem1)
	}

	_, err = cache.GetAndDeleteItem("fjfjf")

	if err == nil && err != ErrNoItem {
		t.Errorf("Expected ErrNoItem")
	}

	_, err = (&Cache{}).GetAndDeleteItem("fjfjf")

	if err != ErrEmptyCache {
		t.Errorf("Expected ErrEmptyCache, got %e", err)
	}
}

func TestCache_SetItem(t *testing.T) {
	cache := initTestingCache()

	cache.SetItem("key2", 2)

	item, found := cache.items["key2"]

	if !found {
		t.Errorf("Expected %d, instead not found", TestingItem2)
	} else if item != TestingItem2 {
		t.Errorf("Expected %d, got %d", TestingItem2, item)
	}
}

func TestCache_Drop(t *testing.T) {
	cache := initTestingCache()

	cache.Drop()

	if len(cache.items) != 0 {
		t.Errorf("Expected zero length")
	}
}
