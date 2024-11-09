package des1

// GetSubList returns sub list started with start and finished with last element
// It doesn't copy items in the list, changes in sub list can affect the original list.
func GetSubList(start int, list LinkedList) LinkedList {
	if list.IsEmpty() {
		// asc what should we return if list is empty
		return LinkedList{}
	}
	if start == 0 {
		return list
	}
	n := list.head
	i := 1
	for n.Next != nil {
		if i == start {
			return LinkedList{head: n.Next}
		}
		i++
		n = n.Next
	}
	return LinkedList{}
}
