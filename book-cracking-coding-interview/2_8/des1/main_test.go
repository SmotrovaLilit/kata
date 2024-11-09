package des1

import "testing"

func TestLoop(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		ls := CreateFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		end := ls.head
		for end.Next != nil {
			end = end.Next
		}
		end.Next = ls.head
		if got := LoopDetection(ls); got != 1 {
			t.Errorf("LoopDetection() = %v, want %v", got, 1)
		}
	})
	t.Run("loop in middle", func(t *testing.T) {
		ls := CreateFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		end := ls.head
		for end.Next != nil {
			end = end.Next
		}
		end.Next = ls.head.Next
		if got := LoopDetection(ls); got != 2 {
			t.Errorf("LoopDetection() = %v, want %v", got, 2)
		}
	})
	t.Run("loop in the end", func(t *testing.T) {
		ls := CreateFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		end := ls.head
		for end.Next != nil {
			end = end.Next
		}
		end.Next = end
		if got := LoopDetection(ls); got != 9 {
			t.Errorf("LoopDetection() = %v, want %v", got, 9)
		}
	})
}
