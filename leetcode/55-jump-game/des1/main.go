package des1

func canJump(nums []int) bool {
	if len(nums) == 1 && nums[0] == 0 {
		return true
	}
	return canJumpFromPosition(nums, 0)
}

func canJumpFromPosition(nums []int, position int) bool {
	for j := nums[position]; j > 0; j-- {
		newPos := position + j
		if newPos >= len(nums)-1 {
			return true
		}
		if nums[newPos] == 0 {
			continue
		}
		canJump := canJumpFromPosition(nums, newPos)
		if canJump {
			return true
		}
	}
	return false
}
