type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

for _, s := range strs {
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteString("#")
	b.WriteString(s)
}

return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var j int
	out := make([]string, 0)
	
	// 3#abc1##
	for i := 0; i < len(encoded);i++{
		if encoded[i] == '#'{
			strLen, _ := strconv.Atoi(encoded[j:i])
			out = append(out, encoded[i+1:i+1+strLen])
			j = i+1+strLen
			i = j
		}
	}

	return out
}
