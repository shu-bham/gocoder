package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Simple Put and Get", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		cache.Put(1, 1)
		assert.Equal(t, 1, cache.Get(1))
	})

	t.Run("Get non-existent key", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		assert.Equal(t, -1, cache.Get(1))
	})

	t.Run("Eviction", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3) // Evicts key 1
		assert.Equal(t, -1, cache.Get(1))
		assert.Equal(t, 2, cache.Get(2))
		assert.Equal(t, 3, cache.Get(3))
	})

	t.Run("Update value", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		cache.Put(1, 1)
		cache.Put(1, 10)
		assert.Equal(t, 10, cache.Get(1))
	})

	t.Run("Get causes reordering", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		_ = cache.Get(1) // 1 is now most recently used
		cache.Put(3, 3)  // Evicts key 2
		assert.Equal(t, -1, cache.Get(2))
		assert.Equal(t, 1, cache.Get(1))
		assert.Equal(t, 3, cache.Get(3))
	})

	t.Run("Complex scenario", func(t *testing.T) {
		cache := LRUCacheConstructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		assert.Equal(t, 1, cache.Get(1))
		cache.Put(3, 3) // evicts key 2
		assert.Equal(t, -1, cache.Get(2))
		cache.Put(4, 4) // evicts key 1
		assert.Equal(t, -1, cache.Get(1))
		assert.Equal(t, 3, cache.Get(3))
		assert.Equal(t, 4, cache.Get(4))
	})
}
