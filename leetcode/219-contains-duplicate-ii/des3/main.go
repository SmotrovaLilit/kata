package des3

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
		if len(set) > k {
			delete(set, nums[i-k])
		}
	}
	return false
}
