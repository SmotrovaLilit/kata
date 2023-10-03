package des1

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	distinctWaysI := make([]int, n+1)
	distinctWaysI[1] = 1
	distinctWaysI[2] = 2
	for i := 3; i <= n; i++ {
		distinctWaysI[i] = distinctWaysI[i-1] + distinctWaysI[i-2]
	}
	return distinctWaysI[n]
}
