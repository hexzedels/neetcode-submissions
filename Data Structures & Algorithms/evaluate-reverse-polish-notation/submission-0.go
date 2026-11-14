func evalRPN(tokens []string) int {
	stack := []int{}

	var result int

	for _, r := range tokens {
		i, err := strconv.Atoi(r)
		if err != nil {
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch r {
			case "+":
				result = left + right
			case "*":
				result = left * right
			case "-":
				result = left - right
			case "/":
				result = left / right
			}
			stack = append(stack, result)
		} else {
			stack = append(stack, i)
		}
	}

    return stack[0]
}