package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	require.True(t, trie.Search("apple"))
	require.False(t, trie.Search("app"))
	require.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	require.True(t, trie.Search("app"))
}
