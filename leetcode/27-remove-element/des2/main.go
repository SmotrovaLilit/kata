package des1

func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; {
		if nums[i] == val {
			nums[i] = nums[numsLen-1]
			numsLen--
		} else {
			i++
		}
	}
	return numsLen
}
