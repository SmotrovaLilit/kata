package des1

func twoSum(nums []int, target int) []int {
	numsIndexes := make(map[int]int)
	for i, v := range nums {
		j, ok := numsIndexes[target-v]
		if ok {
			return []int{j, i}
		}
		numsIndexes[v] = i
	}
	return nil
}
