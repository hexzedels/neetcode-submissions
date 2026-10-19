func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break // positive integers cannot form triplet
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue // we already counted this triplet
		}

		lo, hi := i+1, n-1

		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]

			if sum < 0 {
				lo++
			} else if sum > 0 {
				hi--
			} else {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}

				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}

				lo++
				hi--
			}
		}
	}

	return res
}