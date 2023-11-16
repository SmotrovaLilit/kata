package calc

import "strings"

var ops = []string{"+", "-", "*", "/"}

func calc(expr string) float64 {
	for i := 0; i < len(ops); i++ {
		arguments := strings.SplitN(expr, ops[i], 2)
		if len(arguments) != 2 {
			panic("invalid expression")
		}
		left, right := arguments[0], arguments[1]
		switch ops[i] {
		case "+":
			return calc(left) + calc(right)
		case "-":
			return calc(left) - calc(right)
		case "*":
			return calc(left) * calc(right)
		case "/":
			return calc(left) / calc(right)
		}

	}
	return 0
}
