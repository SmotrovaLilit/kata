package des1

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		i := start + (end-start)/2
		if nums[i] == target {
			return i
		}
		if target < nums[i] {
			end = i - 1
			continue
		}
		start = i + 1

	}
	return -1
}
