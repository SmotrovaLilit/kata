package des1

func twoSum(nums []int, target int) []int {
	var v2 int
	for i, v1 := range nums {
		v2 = target - v1
		if v2 > nums[len(nums)-1] {
			continue
		}
		if v2 < nums[i+1] {
			continue
		}
		if j := search(nums[i+1:], v2); j != -1 {
			return []int{i + 1, i + 1 + j + 1}
		}
	}
	return nil
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
			continue
		}
		end = end - 1
	}
	return -1
}
