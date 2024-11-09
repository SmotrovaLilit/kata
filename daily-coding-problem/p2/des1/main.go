package des1

func MultiplyExceptCurrentIndex(in []int) []int {
	all := 1
	res := make([]int, len(in))
	hasZero := false
	for i := 0; i < len(in); i++ {
		if in[i] != 0 {
			all = all * in[i]
			continue
		}
		hasZero = true
	}
	for i := 0; i < len(in); i++ {
		if in[i] == 0 {
			res[i] = all
			continue
		}
		if hasZero {
			res[i] = 0
			continue
		}
		res[i] = all / in[i]
	}
	return res
}
