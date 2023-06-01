package des1

func search(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	if target == nums[start] {
		return true
	}
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return true
		}
		if nums[start] == nums[mid] { // we can't know where mid in first or second array
			start++
			continue
		}
		if nums[start] < nums[mid] { // mid lies in first sorted array
			if nums[start] <= target { //  target lies in first sorted array
				if nums[mid] > target {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				start = mid + 1
			}
		} else { // mid lies in second sorted array
			if nums[start] <= target { // target lies in first sorted array
				end = mid - 1
			} else { // target lies in second sorted array
				if nums[mid] > target {
					end = mid - 1
				} else {
					start = mid + 1
				}
			}
		}
	}
	return false
}
