func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	minWindowStart := 0
	minWindowEnd := 0
	minWindow := math.MaxInt32
	lo := 0

	freqT := [84]int{}

	for _, r := range t {
		freqT[r-'A']++
	}

	freqS := [84]int{}

	for hi := 0; hi <= len(s)-1; hi++ {
		freqS[s[hi]-'A']++

		for lo <= hi && contains(freqT, freqS) {
			if hi-lo+1 < minWindow {
				minWindowStart = lo
				minWindowEnd = hi
				minWindow = hi - lo + 1
			}

			freqS[s[lo]-'A']--
			lo++
		}
	}

	if minWindow == math.MaxInt32 {
		return ""
	}

	return s[minWindowStart : minWindowEnd+1]
}

// This is O(1)
func contains(first, second [84]int) bool {
	for i, freq := range first {
		if second[i] < freq {
			return false
		}
	}
	return true
}