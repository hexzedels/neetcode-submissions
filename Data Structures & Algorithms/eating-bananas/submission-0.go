func minEatingSpeed(piles []int, h int) int {
	maxPile := 0

	for _, p := range piles {
		maxPile = max(maxPile, p)
	}

	lo := 1
	hi := maxPile

	k := maxPile

	for lo <= hi {
		mid := lo + (hi-lo)/2

		tempH := 0

		for _, p := range piles {
			tempH += (p + mid - 1) / mid

			if tempH > h {
				lo = mid + 1
				break
			}
		}

		if tempH <= h {
			k = min(k, mid)
			hi = mid - 1
		}
	}

	return k
}