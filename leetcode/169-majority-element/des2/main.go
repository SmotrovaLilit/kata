package des1

func majorityElement(nums []int) int {
	majCount := len(nums) / 2
	mapCount := make(map[int]int)
	for _, v := range nums {
		mapCount[v]++
		if mapCount[v] > majCount {
			return v
		}
	}
	return -1
}
