package des1

import "sort"

func merge(nums1 []int, m int, nums2 []int, _ int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
