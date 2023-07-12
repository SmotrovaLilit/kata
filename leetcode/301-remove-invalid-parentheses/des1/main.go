package des1

import (
	"math"
)

var (
	minimumRemoved   int
	validExpressions map[string]struct{}
)

func removeInvalidParentheses(s string) []string {
	minimumRemoved = math.MaxInt
	validExpressions = make(map[string]struct{})
	remain(s, 0, 0, 0, 0, []byte{})
	result := make([]string, 0, len(validExpressions))
	for k, _ := range validExpressions {
		result = append(result, k)
	}
	return result
}

func remain(s string, index, leftCount, rightCount, removedCount int, expression []byte) []byte {
	if index == len(s) { // If we have reached the end of string.
		if leftCount == rightCount { // If the current expression is valid.
			if removedCount <= minimumRemoved { // If the current count of removed parentheses is <= the current minimum count
				// Convert StringBuilder to a String. This is an expensive operation.
				// So we only perform this when needed.
				possibleAnswer := string(expression)
				// If the current count beats the overall minimum we have till now
				if removedCount < minimumRemoved {
					validExpressions = make(map[string]struct{})
					minimumRemoved = removedCount
				}
				validExpressions[possibleAnswer] = struct{}{}
			}
		}
	} else {
		currentChar := s[index]
		length := len(expression)
		// If the current character is neither an opening bracket nor a closing one,
		// simply recurse further by adding it to the expression
		if currentChar != '(' && currentChar != ')' {
			expression = append(expression, currentChar)
			expression = remain(s, index+1, leftCount, rightCount, removedCount, expression)

			// Remove currentChar from expression
			// Undoing the append operation for other recursions.
			expression = append(expression[:length], expression[length+1:]...)
		} else {
			// Recursion where we delete the current character and move forward
			expression = remain(s, index+1, leftCount, rightCount, removedCount+1, expression)
			expression = append(expression, currentChar)

			// If it's an opening parenthesis, consider it and recurse
			if currentChar == '(' {
				expression = remain(s, index+1, leftCount+1, rightCount, removedCount, expression)
			} else if rightCount < leftCount {
				// For a closing parenthesis, only recurse if right < left
				expression = remain(s, index+1, leftCount, rightCount+1, removedCount, expression)
			}
			// Remove currentChar from expression
			// Undoing the append operation for other recursions.
			expression = append(expression[:length], expression[length+1:]...)
		}
	}
	return expression
}
