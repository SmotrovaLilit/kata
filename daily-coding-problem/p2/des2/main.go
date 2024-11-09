package des2

func MultiplyExceptCurrentIndex(in []int) []int {
	res := make([]int, len(in))
	previousProduct := 1 // product for numbers before i
	productExceptI := 1  // product for numbers except i
	for i := 0; i < len(in); i++ {
		productExceptI = previousProduct
		for j := i + 1; j < len(in); j++ {
			productExceptI = productExceptI * in[j]
		}
		res[i] = productExceptI
		previousProduct = previousProduct * in[i]
	}

	return res
}
