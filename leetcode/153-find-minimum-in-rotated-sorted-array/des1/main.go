package des1

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	if nums[0] <= nums[end] {
		return nums[0]
	}
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] >= nums[start] {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return -1
}
