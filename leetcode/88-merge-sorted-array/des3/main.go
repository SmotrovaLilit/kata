package des1

func merge(nums1 []int, m int, nums2 []int, n int) {
	iNums1, iNums2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if iNums2 < 0 {
			break
		}
		if iNums1 >= 0 && nums1[iNums1] > nums2[iNums2] {
			nums1[i] = nums1[iNums1]
			iNums1--
			continue
		}
		nums1[i] = nums2[iNums2]
		iNums2--
	}
}
