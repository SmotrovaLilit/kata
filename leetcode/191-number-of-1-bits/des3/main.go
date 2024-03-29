package des3

func hammingWeight(num uint32) int {
	count := 0
	for i := uint32(0); i <= 32; i++ {
		if num&(1<<i) != 0 {
			count++
		}
	}
	return count
}
