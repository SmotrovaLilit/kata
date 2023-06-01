package des1

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[start] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
				continue
			}
			start = mid + 1
			continue
		}
		if target <= nums[end] && target > nums[mid] {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return -1
}
