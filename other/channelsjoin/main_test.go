package channelsjoin

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_joinChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	got := joinChannels(ch1, ch2)
	go func() {
		ch1 <- 1
		ch2 <- 2
		close(ch1)
		ch2 <- 3
		ch2 <- 4
		time.Sleep(5 * time.Second)
		close(ch2)
	}()
	expectedValues := []int{1, 2, 3, 4}
	gotValues := make([]int, 0, len(expectedValues))
	for v := range got {
		gotValues = append(gotValues, v)
	}
	require.ElementsMatch(t, expectedValues, gotValues)
}
