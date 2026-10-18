func isValidSudoku(board [][]byte) bool {
	var rows [9][9]uint8
	var columns [9][9]uint8
	var subbox [3][3][9]uint8

	subMap := map[int]int{0: 0, 1: 0, 2: 0, 3: 1, 4: 1, 5: 1, 6: 2, 7: 2, 8: 2}

	for i, row := range board {
		for j, v := range row {
			if rune(v) != '.' {
				r, _ := strconv.Atoi(string(v))
				r--
				if rows[i][r] == 1 {
					return false
				} else {
					rows[i][r]++
				}
				if columns[j][r] == 1 {
					return false
				} else {
					columns[j][r]++
				}

				x := subMap[i]
				y := subMap[j]

				if subbox[x][y][r] == 1 {
					return false
				} else {
					subbox[x][y][r]++
				}
			}
		}
	}
	return true
}