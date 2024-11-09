package biginteger

import "sort"

func sum(num1, num2 []int) []int {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := make([]int, len(num1)+1)
	resTmp := res[1:]

	c := 0
	for j, k := len(num2)-1, len(num1)-1; k >= 0; j, k = j-1, k-1 {
		resTmp[k] = num1[k]
		if j >= 0 {
			resTmp[k] += num2[j]
		}
		resTmp[k] += c
		c = resTmp[k] / 10
		resTmp[k] %= 10
	}
	if c == 0 {
		return resTmp
	}
	res[0] = c
	return res
}

func ff2(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys[:k]
}

func fff(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	keys := nums[0:len(m)]
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys[:k]
}
