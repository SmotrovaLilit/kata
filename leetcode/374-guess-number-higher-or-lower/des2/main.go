package des1

var (
	testGuess = func(pick int) func(num int) int {
		return func(num int) int {
			if num > pick {
				return -1
			} else if num < pick {
				return 1
			} else {
				return 0
			}
		}
	}
	guess = testGuess(5)
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;.
 */
func guessNumber(n int) int {
	start := 0
	end := n
	var m1, m2, guessResult1, guessResult2 int
	for start <= end {
		m1 = start + (end-start)/3
		m2 = end - (end-start)/3
		guessResult1 = guess(m1)
		if guessResult1 == 0 {
			return m1
		}
		guessResult2 = guess(m2)
		if guessResult2 == 0 {
			return m2
		}
		if guessResult1 == -1 { // pick < m1
			end = m1 - 1
			continue
		}
		if guessResult2 == 1 { // pick > m2
			start = m2 + 1
			continue
		}
		// m1 < pick < m2
		start = m1 + 1
		end = m2 - 1
		continue
	}
	return -1
}
