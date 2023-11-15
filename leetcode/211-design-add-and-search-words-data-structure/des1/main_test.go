package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	require.False(t, wordDictionary.Search("pad"))
	require.True(t, wordDictionary.Search("bad"))
	require.True(t, wordDictionary.Search(".ad"))
	require.True(t, wordDictionary.Search("b.."))
}
