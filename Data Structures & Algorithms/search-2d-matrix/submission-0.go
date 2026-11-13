func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix[0])

	rowLo, rowHi := 0, len(matrix)-1

	for rowLo <= rowHi {
		mid := rowLo + (rowHi-rowLo)/2

		switch {
		case matrix[mid][m-1] >= target && matrix[mid][0] <= target:
			colLo, colHi := 0, m-1
			for colLo <= colHi {
				colMid := colLo + (colHi-colLo)/2

				switch {
				case matrix[mid][colMid] == target:
					return true
				case matrix[mid][colMid] > target:
					colHi = colMid - 1
				default:
					colLo = colMid + 1
				}
			}
            return false
		case matrix[mid][0] > target:
			rowHi = mid - 1
		case matrix[mid][m-1] < target:
			rowLo = mid + 1
		}
	}

	return false
}