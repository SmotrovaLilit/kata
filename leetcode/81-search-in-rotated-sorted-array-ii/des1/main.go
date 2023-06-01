package des1

func search(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return true
		}
		if start != end && nums[start] == nums[mid] && nums[mid] == nums[end] {
			// need simple searching
			for i := start; i <= end; i++ {
				if nums[i] != nums[start] {
					if nums[i] == target {
						return true
					}
					return false
				}
			}
			return false
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
	return false
}
