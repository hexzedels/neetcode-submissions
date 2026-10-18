type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var out strings.Builder
	del := "#"

	for _, str := range strs {
		out.WriteString(strconv.Itoa(len(str)))
		out.WriteString(del)
		out.WriteString(str)
	}

	return out.String()
}

func (s *Solution) Decode(encoded string) []string {
	var answer []string
	var l int
	var j int

	for i := 0; i < len(encoded);i++{
		if encoded[i] == '#' {
			l, _ = strconv.Atoi(string(encoded[j:i]))
			answer = append(answer, string(encoded[i+1:i+1+l]))
			i = i+l
			j = i + 1
		}
	}

	return answer
}
