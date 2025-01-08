package time_based_key_value_store_981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKV(t *testing.T) {
	kv := Constructor()
	kv.Set("foo", "bar", 1)
	kv.Set("foo", "baz", 4)
	assert.Equal(t, "", kv.Get("foo", 0))
	assert.Equal(t, "bar", kv.Get("foo", 3))
	assert.Equal(t, "bar", kv.Get("foo", 2))
	assert.Equal(t, "baz", kv.Get("foo", 4))
	assert.Equal(t, "baz", kv.Get("foo", 6))
}
