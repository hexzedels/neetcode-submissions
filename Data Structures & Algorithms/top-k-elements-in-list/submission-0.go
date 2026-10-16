func topKFrequent(nums []int, k int) []int {
	maxFreq := 0
	freq := make(map[int]int, 0)
	for _, n := range nums {
		nFreq := freq[n] + 1
		freq[n] = nFreq
		if nFreq > maxFreq {
			maxFreq = nFreq
		}
	}

	sorted := make([][]int, maxFreq)

	for value, fr := range freq {
		sorted[fr-1] = append(sorted[fr-1], value)
	}

	answ := make([]int, 0, k)
	hi := maxFreq - 1
	for k > 0 {
		for _, v := range sorted[hi] {
			answ = append(answ, v)
			k--
			if k == 0 {
				return answ
			}
		}
		hi--
	}
	return answ
}
