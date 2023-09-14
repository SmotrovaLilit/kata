package des1

import (
	"sort"
)

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	stackRandom := newStack()
	stackMagazine := newStack()
	ransomBytes := []byte(ransomNote)
	magazineBytes := []byte(magazine)
	sort.Slice(ransomBytes, func(i, j int) bool {
		return ransomBytes[i] > ransomBytes[j]
	})
	sort.Slice(magazineBytes, func(i, j int) bool {
		return magazineBytes[i] > magazineBytes[j]
	})

	for i := 0; i < len(ransomBytes); i++ {
		stackRandom.push(ransomBytes[i])
	}
	for i := 0; i < len(magazineBytes); i++ {
		stackMagazine.push(magazineBytes[i])
	}
	for !stackRandom.isEmpty() && !stackMagazine.isEmpty() {
		if stackRandom.peek() == stackMagazine.peek() {
			stackRandom.pop()
			stackMagazine.pop()
		} else if stackRandom.peek() > stackMagazine.peek() {
			stackMagazine.pop()
		} else {
			return false
		}
	}
	return stackRandom.isEmpty()
}

type listNode struct {
	val  byte
	next *listNode
}

type stack struct {
	top *listNode
}

func newStack() stack {
	return stack{}
}

func (s *stack) push(b byte) {
	s.top = &listNode{
		val:  b,
		next: s.top,
	}
}

func (s *stack) pop() byte {
	if s.isEmpty() {
		panic("stack is empty")
	}
	v := s.top.val
	s.top = s.top.next
	return v
}

func (s *stack) peek() byte {
	if s.isEmpty() {
		panic("stack is empty")
	}
	return s.top.val
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}
