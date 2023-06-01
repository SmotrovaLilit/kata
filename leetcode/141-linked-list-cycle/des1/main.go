package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	current := head
	nodePointers := map[*ListNode]struct{}{}
	for {
		if current == nil {
			return false
		}
		if _, ok := nodePointers[current]; ok {
			return true
		}
		nodePointers[current] = struct{}{}
		current = current.Next
	}
}
