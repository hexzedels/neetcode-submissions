type Solution struct{}

func (s *Solution) Encode(strs []string) string {
var sb strings.Builder

	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteString("#")
		sb.WriteString(s)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string

var start int
	for i := 0; i < len(encoded); i++ {
		if encoded[i] == '#' {
			l, _ := strconv.Atoi(encoded[start:i])
			result = append(result, encoded[i+1:i+1+l])
			i = i+1+l
			start = i
		}
	}

	return result
}
