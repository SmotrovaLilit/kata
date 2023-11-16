package fib

// fib returns the nth Fibonacci number.
// Time complexity: O(n).
// Space complexity: O(n).
func fib(n int) int {
	cache := make(map[int]int)
	f := func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		if n < 0 {
			panic("n must be greater than or equal to 0")
		}
		if n <= 1 {
			return 1
		}

		r := fib(n-1) + fib(n-2)
		cache[n] = r
		return r
	}
	return f(n)
}
