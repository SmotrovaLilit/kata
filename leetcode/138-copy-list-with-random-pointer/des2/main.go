package des2

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	oldNode := head
	var p *Node
	// create  A->A'->B->B'->C->C'
	for oldNode != nil {
		p = oldNode.Next
		oldNode.Next = &Node{
			Val:  oldNode.Val,
			Next: p,
		}
		oldNode = p
	}
	oldNode = head
	// change random links in copies to copies
	for oldNode != nil {
		if oldNode.Random != nil {
			p = oldNode.Next
			p.Random = oldNode.Random.Next
		}
		oldNode = oldNode.Next.Next
	}
	// split list i.e. A->A'->B->B'->C->C' would be broken to A->B->C and A'->B'->C'
	oldNode = head
	newNode := head.Next
	headNew := head.Next
	for oldNode != nil {
		oldNode.Next = oldNode.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}

		newNode = newNode.Next
		oldNode = oldNode.Next
	}
	return headNew
}
