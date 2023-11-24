package des1

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
	newNodesMap := make(map[*Node]*Node)
	newNode := &Node{Val: head.Val}
	newNodesMap[head] = newNode

	for oldNode != nil {
		if v, ok := newNodesMap[oldNode.Next]; ok {
			newNode.Next = v
		} else if oldNode.Next != nil {
			newNode.Next = &Node{
				Val: oldNode.Next.Val,
			}
			newNodesMap[oldNode.Next] = newNode.Next
		}

		if v, ok := newNodesMap[oldNode.Random]; ok {
			newNode.Random = v
		} else if oldNode.Random != nil {
			newNode.Random = &Node{
				Val: oldNode.Random.Val,
			}
			newNodesMap[oldNode.Random] = newNode.Random
		}

		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return newNodesMap[head]
}
