package des1

func singleNumber(nums []int) int {
	xorResult := 0
	for _, v := range nums {
		xorResult = xorResult ^ v
	}
	return xorResult
}
