package des1

func canCompleteCircuit(gas []int, cost []int) int {
	currentGain := 0
	totalGain := 0
	answer := 0
	for i := 0; i < len(gas); i++ {
		currentGain += gas[i] - cost[i]
		totalGain += gas[i] - cost[i]
		if currentGain < 0 {
			answer = i + 1
			currentGain = 0
		}
	}
	if totalGain < 0 {
		return -1
	}
	return answer
}
