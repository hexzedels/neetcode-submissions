func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freqS1 := [26]int{}

	for _, r := range s1 {
		freqS1[r-'a']++
	}

	width := len(s1)

	lo := 0


freqS2 := [26]int{}

	for hi, r := range s2 {
		freqS2[r-'a']++
		
		if hi - lo + 1 > width {
			freqS2[s2[lo]-'a']--
			lo++
		}

		if freqS1 == freqS2 {
			return true
		}
	}

	return false
}
