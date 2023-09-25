package des1

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		rightMostBit := num & 1 // retrieve right-most bit of an integer
		ret += rightMostBit << power
		num = num >> 1
		power -= 1
	}
	return ret
}
