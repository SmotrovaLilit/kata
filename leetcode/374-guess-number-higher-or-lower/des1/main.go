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
	var i, guessResult int
	for start <= end {
		i = start + (end-start)/2
		guessResult = guess(i)
		if guessResult == 0 {
			return i
		}
		if guessResult == -1 {
			end = i - 1
			continue
		}
		if guessResult == 1 {
			start = i + 1
			continue
		}
	}
	return -1
}
