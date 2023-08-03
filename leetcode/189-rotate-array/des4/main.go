package des1

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	var temp int
	start := 0
	end := len(nums) - 1
	for {
		if start >= end {
			break
		}
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
