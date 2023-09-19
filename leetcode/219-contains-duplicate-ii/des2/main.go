package des2

func containsNearbyDuplicate(nums []int, k int) bool {
	set := newTreeSet()
	for i := 0; i < len(nums); i++ {
		if set.contains(nums[i]) {
			return true
		}
		set.add(nums[i])
		if set.size() > k {
			set.remove(nums[i-k])
		}
	}
	return false
}

type myTreeNode struct {
	Value       int
	Left, Right *myTreeNode
}

type treeSet struct {
	Root *myTreeNode
	Size int
}

func (s *treeSet) contains(i int) bool {
	return containsRecursive(s.Root, i)
}

func (s *treeSet) add(i int) {
	s.Root, _ = insertRecursive(s.Root, i)
	s.Size++
}

func (s *treeSet) size() int {
	return s.Size
}

func (s *treeSet) remove(i int) {
	var removed bool
	s.Root, removed = removeRecursive(s.Root, i)
	if removed {
		s.Size--
	}
}

func containsRecursive(node *myTreeNode, i int) bool {
	if node == nil {
		return false
	}
	if i == node.Value {
		return true
	}
	if i < node.Value {
		return containsRecursive(node.Left, i)
	} else {
		return containsRecursive(node.Right, i)
	}
}
func insertRecursive(node *myTreeNode, i int) (*myTreeNode, bool) {
	if node == nil {
		return &myTreeNode{Value: i}, true
	}
	if i == node.Value {
		return node, false
	}
	if i < node.Value {
		node.Left, _ = insertRecursive(node.Left, i)
	} else {
		node.Right, _ = insertRecursive(node.Right, i)
	}
	return node, true
}

func removeRecursive(node *myTreeNode, i int) (*myTreeNode, bool) {
	if node == nil {
		return nil, false // Не найдено
	}
	if i == node.Value {
		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}
		minValue := findMinValue(node.Right)
		node.Value = minValue
		node.Right, _ = removeRecursive(node.Right, minValue)
		return node, true
	} else if i < node.Value {
		node.Left, _ = removeRecursive(node.Left, i)
	} else {
		node.Right, _ = removeRecursive(node.Right, i)
	}
	return node, true
}

func findMinValue(node *myTreeNode) int {
	if node.Left == nil {
		return node.Value
	}
	return findMinValue(node.Left)
}

func newTreeSet() treeSet {
	return treeSet{}
}
