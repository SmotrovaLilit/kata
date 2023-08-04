package des1

func canJump(nums []int) bool {
	isNotFailedPositions := make([]bool, len(nums))
	isNotFailedPositions[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		fJump := i + nums[i]
		if fJump > len(nums) {
			fJump = len(nums) - 1
		}
		for j := i + 1; j <= fJump; j++ {
			if isNotFailedPositions[j] {
				isNotFailedPositions[i] = true
				break
			}
		}
	}
	return isNotFailedPositions[0]
}
