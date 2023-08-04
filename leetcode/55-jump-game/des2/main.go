package des1

var isFailedPositions []bool

func canJump(nums []int) bool {
	isFailedPositions = make([]bool, len(nums))
	return canJumpFromPosition(nums, 0)
}

func canJumpFromPosition(nums []int, position int) bool {
	if position >= len(nums)-1 {
		return true
	}
	if nums[position] == 0 {
		return false
	}
	if isFailedPositions[position] {
		return false
	}

	for j := nums[position]; j > 0; j-- {
		newPos := position + j
		canJump := canJumpFromPosition(nums, newPos)
		if canJump {
			return true
		}
	}
	isFailedPositions[position] = true
	return false
}
