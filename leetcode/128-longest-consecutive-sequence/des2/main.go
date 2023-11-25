package des1

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	longest := 0
	for num, _ := range m {
		if m[num-1] {
			continue
		}
		end := num
		for m[end+1] {
			end++
		}
		if end-num+1 > longest {
			longest = end - num + 1
		}

	}

	return longest
}
