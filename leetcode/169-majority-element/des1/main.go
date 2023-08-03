package des1

func majorityElement(nums []int) int {
	majCount := len(nums) / 2
	for _, v := range nums {
		count := 0
		for _, v2 := range nums {
			if v == v2 {
				count++
			}
		}
		if count > majCount {
			return v
		}
	}
	return -1
}
