func isValid(s string) bool {
    pair := map[rune]rune{')':'(', ']':'[', '}':'{'}
	st := []rune{}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			st = append(st, r)
		case ')', '}', ']':
			if len(st) == 0 {
				return false
			}
			p := st[len(st)-1]
			st = st[:len(st)-1]

			if p != pair[r] {
				return false
			}
		}
	}

	return len(st) == 0
}