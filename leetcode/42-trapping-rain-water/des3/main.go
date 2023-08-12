package des1

type stack struct {
	data []int
}

func newStack(capacity int) *stack {
	return &stack{data: make([]int, 0, capacity)}
}

func (s *stack) top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) pop() int {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *stack) push(v int) {
	s.data = append(s.data, v)
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

func trap(height []int) int {
	s := newStack(len(height))
	water := 0
	current := 0
	for current < len(height) {
		for !s.empty() && height[current] > height[s.top()] {
			top := s.pop()
			if s.empty() {
				break
			}
			distance := current - s.top() - 1
			boundedHeight := height[current]
			if height[current] > height[s.top()] {
				boundedHeight = height[s.top()]
			}
			boundedHeight = boundedHeight - height[top]
			water += distance * boundedHeight
		}
		s.push(current)
		current++
	}
	return water
}
