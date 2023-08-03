package des1

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[k-1] {
			continue
		}
		nums[k] = nums[i]
		k++
	}
	return k
}
