func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1

	for lo < hi {
		if numbers[lo]+numbers[hi] == target {
			break
		} else if numbers[lo]+numbers[hi] > target {
			hi--
		} else {
			lo++
		}
	}
	return []int{lo + 1, hi + 1}
}