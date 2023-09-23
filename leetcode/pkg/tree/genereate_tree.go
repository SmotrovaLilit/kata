package tree

func GenerateTree(nodeCount int) *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	q := &queue{}
	q.add(root)
	for i := 2; i <= nodeCount; {
		cur := q.poll()
		cur.Left = &TreeNode{
			Val: i,
		}
		q.add(cur.Left)
		i++
		if i > nodeCount {
			break
		}
		cur.Right = &TreeNode{
			Val: i,
		}
		i++
		q.add(cur.Right)
	}
	return root
}

type queueNode struct {
	Val  *TreeNode
	Next *queueNode
}
type queue struct {
	head *queueNode
	tail *queueNode
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}

func (q *queue) add(val *TreeNode) {
	if q.head == nil {
		q.head = &queueNode{
			Val: val,
		}
		q.tail = q.head
		return
	}
	q.tail.Next = &queueNode{
		Val: val,
	}
	q.tail = q.tail.Next

}

func (q *queue) poll() *TreeNode {
	if q.isEmpty() {
		panic("queue is empty")
	}
	val := q.head.Val
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	return val
}
