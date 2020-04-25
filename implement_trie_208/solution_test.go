package implement_trie_208

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	type operation struct {
		method string
		value  string
		want   bool
	}

	testCases := []struct {
		name string
		ops  []operation
	}{
		{
			name: "trie test",
			ops: []operation{
				{
					method: "Insert",
					value:  "apple",
				},
				{
					method: "Search",
					value:  "apple",
					want:   true,
				},
				{
					method: "Search",
					value:  "app",
					want:   false,
				},
				{
					method: "StartsWith",
					value:  "app",
					want:   true,
				},
				{
					method: "Insert",
					value:  "app",
				},
				{
					method: "Search",
					value:  "app",
					want:   true,
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, op := range tt.ops {
				switch op.method {
				case "Insert":
					trie.Insert(op.value)
				case "Search":
					assert.Equal(t, op.want, trie.Search(op.value))
				case "StartsWith":
					assert.Equal(t, op.want, trie.StartsWith(op.value))
				}
			}
		})
	}
}
