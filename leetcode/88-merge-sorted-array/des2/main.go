package des1

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])
	iNums1, iNums2 := 0, 0
	for i := 0; i < m+n; i++ {
		if iNums1 >= m {
			nums1 = append(nums1[:i], nums2[iNums2:]...)
			return
		}
		if iNums2 >= n {
			nums1 = append(nums1[:i], nums1Copy[iNums1:]...)
			return
		}
		//if iNums1 >= m {
		//	nums1[i] = nums2[iNums2]
		//	iNums2++
		//	continue
		//}
		//if iNums2 >= n {
		//	nums1[i] = nums1Copy[iNums1]
		//	iNums1++
		//	continue
		//}
		if nums1Copy[iNums1] <= nums2[iNums2] {
			nums1[i] = nums1Copy[iNums1]
			iNums1++
			continue
		}
		nums1[i] = nums2[iNums2]
		iNums2++
	}
}
