package des1

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1
	var sum int
	for start < end {
		sum = nums[start] + nums[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum < target {
			start++
			continue
		}
		end--
	}
	return nil
}
