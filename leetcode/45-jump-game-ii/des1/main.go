package des1

func jump(nums []int) int {
	answer := 0
	curEnd := 0
	curFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i + nums[i]
		}
		if i == curEnd {
			answer++
			curEnd = curFarthest
		}
	}
	return answer
}
