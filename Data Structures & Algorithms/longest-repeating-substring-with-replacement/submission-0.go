func characterReplacement(s string, k int) int {
	freq := [26]int{}

	lo := 0

	best := 0
	for hi, r := range s {
		freq[r-'A']++

		// Move lo to right while we in substring we cannot repair
		for lo < hi && hi-lo+1-maxFreq(&freq) > k {
			freq[s[lo]-'A']--
			lo++
		}

		best = max(best, hi-lo+1)
	}

	return best
}

func maxFreq(freq *[26]int) int {
	mFr := 0

	for _, f := range *freq {
		mFr = max(mFr, f)
	}

	return mFr
}