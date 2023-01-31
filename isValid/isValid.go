// lc 20
package main

import "log"

func main() {
	s := "]"

	log.Println(isValid(s))
}

// '('，')'，'{'，'}'，'['，']'
func isValid2(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	stack := make([]int32, 0, len(s))
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')':
			if len(stack) < 1 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) < 1 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ']':
			if len(stack) < 1 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) < 1 {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := make([]int32, 0)

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append([]int32{c} , stack...)
		case ')':
			if len(stack) > 0 && stack[0] == '(' {
				stack = stack[1:]
			}else{
				return false
			}
		case '}':
			if len(stack) > 0 && stack[0] == '{' {
				stack = stack[1:]
			}else{
				return false
			}
		case ']':
			if len(stack) > 0 && stack[0] == '[' {
				stack = stack[1:]
			}else{
				return false
			}
		default:
			break
		}
	}
	return len(stack) == 0
}
