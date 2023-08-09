package des1

func candy(ratings []int) int {
	leftToRight := make([]int, len(ratings))
	rightToLeft := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		leftToRight[i] = 1
		rightToLeft[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			leftToRight[i] = leftToRight[i-1] + 1
		}
	}
	sum := leftToRight[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightToLeft[i] = rightToLeft[i+1] + 1
		}
		if leftToRight[i] >= rightToLeft[i] {
			sum += leftToRight[i]
		} else {
			sum += rightToLeft[i]
		}
	}
	return sum
}
