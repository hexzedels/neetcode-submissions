func dailyTemperatures(temperatures []int) []int {
	var stack []t
	result := make([]int, len(temperatures))

	for i, temp := range temperatures {
		if len(stack) > 0 {
				for len(stack)> 0 &&  stack[len(stack)-1].temp < temp {
					prev := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					result[prev.ind] = i - prev.ind
				}
			}
		stack = append(stack, t{temp: temp, ind: i})
	}
	return result
}

type t struct {
	temp int
	ind int
}
