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
	result := make([]string, 0)

	start := 0

	for i := 0; i < len(encoded);i++{
		if encoded[i] == '#' {
			ln, _ := strconv.Atoi(encoded[start:i])
			result = append(result, encoded[i+1:i+1+ln])
			start = i+1+ln
			i += ln
		}
	}

	return result
}
