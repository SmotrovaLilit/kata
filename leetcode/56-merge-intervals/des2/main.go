package des1

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	graph := newGraph()
	for i := 0; i < len(intervals); i++ {
		if !graph.existVertex(i) {
			graph.addVertex(i)
		}
		for j := i + 1; j < len(intervals); j++ {
			if !graph.existVertex(j) {
				graph.addVertex(j)
			}
			if isOverlapped(intervals[i], intervals[j]) {
				graph.addEdge(i, j)
			}
		}
	}
	result := make([][]int, 0)
	visited := make(map[int]struct{})
	for i := range graph.Vertices {
		if _, ok := visited[i]; ok {
			continue
		}
		stack := newStack()
		stack.push(i)
		var newIntervalSet [][]int
		for !stack.isEmpty() {
			v := stack.pop()
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = struct{}{}
			newIntervalSet = append(newIntervalSet, intervals[v])
			for _, w := range graph.Vertices[v] { // add to stack other vertices connected to v
				if _, ok := visited[w]; ok {
					continue
				}
				stack.push(w)
			}
		}
		result = append(result, mergeOverlappedIntervals(newIntervalSet))
	}
	return result
}

type node struct {
	value int
	next  *node
}
type stack struct {
	top *node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(value int) {
	s.top = &node{value, s.top}
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type graph struct {
	Vertices map[int][]int
}

func newGraph() *graph {
	return &graph{
		Vertices: make(map[int][]int),
	}
}

func (g *graph) addVertex(v int) {
	if _, ok := g.Vertices[v]; !ok {
		g.Vertices[v] = []int{}
	}
}

func (g *graph) existVertex(v int) bool {
	_, ok := g.Vertices[v]
	return ok
}

func (g *graph) addEdge(v1, v2 int) {
	g.Vertices[v1] = append(g.Vertices[v1], v2)
	g.Vertices[v2] = append(g.Vertices[v2], v1)
}

func mergeOverlappedIntervals(intervals [][]int) []int {
	x, y := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		x = min(x, intervals[i][0])
		y = max(y, intervals[i][1])
	}
	return []int{x, y}
}

func isOverlapped(a []int, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}
