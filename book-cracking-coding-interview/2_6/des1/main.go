package des1

func Palindrome(l LinkedList) bool {
	reversed := l.Reverse()
	n1 := l.head
	n2 := reversed.head
	for n1 != nil && n2 != nil {
		if n1.Value != n2.Value {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return true
}
