package mergingnchannels

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Merge(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	got := Merge(ch1, ch2)
	go func() {
		ch1 <- "1"
		ch2 <- "2"
		close(ch1)
		ch2 <- "3"
		ch2 <- "4"
		time.Sleep(time.Second)
		close(ch2)
	}()
	expectedValues := []string{"1", "2", "3", "4"}
	gotValues := make([]string, 0, len(expectedValues))
	for v := range got {
		gotValues = append(gotValues, v)
	}
	require.ElementsMatch(t, expectedValues, gotValues)
}
