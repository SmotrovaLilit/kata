package des1

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	var triplet []int
	seen := make(map[int]int)
	duplicates := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := duplicates[nums[i]]; ok {
			continue
		}
		duplicates[nums[i]] = struct{}{}
		for j := i + 1; j < len(nums); j++ {
			complement := 0 - nums[i] - nums[j]
			it, ok := seen[complement]
			if ok && it == i {
				triplet = []int{nums[i], nums[j], complement}
				sort.Ints(triplet)
				if !existTriplet(result, triplet) {
					result = append(result, triplet)
				}
			}
			seen[nums[j]] = i // убеждаемся что мы нашли элемент в той же итерации i

		}
	}
	return result
}

func existTriplet(result [][]int, triplet []int) bool {
	for _, list := range result {
		if list[0] == triplet[0] && list[1] == triplet[1] && list[2] == triplet[2] {
			return true
		}
	}
	return false
}
