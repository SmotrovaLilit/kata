package des1

func RemoveDuplicates(list LinkedList) {
	if list.IsEmpty() {
		return
	}
	mapList := make(map[int]struct{})
	n := list.head
	mapList[n.Value] = struct{}{}
	for n.Next != nil {
		if _, ok := mapList[n.Next.Value]; ok {
			n.Next = n.Next.Next
		} else {
			mapList[n.Next.Value] = struct{}{}
			n = n.Next
		}
	}
}
