package lru_cache_146

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache_1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1)) // returns 1

	cache.Put(3, 3)                   // evicts key 2
	assert.Equal(t, -1, cache.Get(2)) // returns -1 (not found)

	cache.Put(4, 4)                   // evicts key 1
	assert.Equal(t, -1, cache.Get(1)) // returns -1 (not found)
	assert.Equal(t, 3, cache.Get(3))  // returns 3
	assert.Equal(t, 4, cache.Get(4))  // returns 4
}

func TestLRUCache_2(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	assert.Equal(t, 1, cache.Get(2))
}

/*
["LRUCache","put","get","put","get","get"]
[[1],[2,1],[2],[3,2],[2],[3]]
*/
func TestLRUCache_3(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	assert.Equal(t, 1, cache.Get(2))

	cache.Put(3, 2)
	assert.Equal(t, -1, cache.Get(2))
	assert.Equal(t, 2, cache.Get(3))
}

/*
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
*/
func TestLRUCache_4(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, 2, cache.Get(2))
}

/*
["LRUCache","put","put","put","put","get","get"]
[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
Output:   [null,null,null,null,null,1,-1]
Expected: [null,null,null,null,null,-1,3]
*/
func TestLRUCache_5(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3) // evict 2->1
	cache.Put(4, 1) // evict 1->1

	// left with: 2->3, 4->1
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(2))
}

/*
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
Output:   [null,null,null,null,null,1,-1]
Expected: [null,null,null,null,null,-1,3]
*/
func TestLRUCache_6(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, 2, cache.Get(2))
	cache.Put(1, 1)
	cache.Put(4, 1)
	assert.Equal(t, -1, cache.Get(2))
}

/*
["LRUCache","get","put","get","put","put","get","get"]
[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
*/
func TestLRUCache_7(t *testing.T) {
	cache := Constructor(2)
	assert.Equal(t, -1, cache.Get(2))
	cache.Put(2, 6)
	assert.Equal(t, -1, cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2) // evicts 2->6
	assert.Equal(t, 2, cache.Get(1))
	assert.Equal(t, 6, cache.Get(2))
}

/*
["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
Output: [null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
Expected: [null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
*/
func TestLRUCache_8(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	assert.Equal(t, 4, cache.Get(4))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 2, cache.Get(2))
	assert.Equal(t, -1, cache.Get(1))
	cache.Put(5, 5)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 2, cache.Get(2))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, -1, cache.Get(4))
	assert.Equal(t, 5, cache.Get(5))
}
