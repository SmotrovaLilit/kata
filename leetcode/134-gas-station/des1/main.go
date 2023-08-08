package des1

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		if canCompleteFromPosition(gas, cost, i) {
			return i
		}
	}
	return -1
}

func canCompleteFromPosition(gas []int, cost []int, position int) bool {
	tank := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[position] - cost[position]
		if tank < 0 {
			return false
		}
		position++
		if position == len(gas) {
			position = 0
		}
	}
	return true
}
