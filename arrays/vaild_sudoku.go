package arrays

func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]bool{} // Index of the row in the board mapped to a set of numbers
	cols := [9]map[byte]bool{} // Index of the cols in the board mapped to a set of numbers
	grid := [9]map[byte]bool{} // Index of the 3x3 grids in the board mapped

	for r, row := range board {
		rows[r] = map[byte]bool{}
		for c, n := range row {
			if cols[c] == nil {
				cols[c] = map[byte]bool{}
			}
			if grid[r/3*3+c/3] == nil {
				grid[r/3*3+c/3] = map[byte]bool{}
			}
			if n != '.' {
				if rows[r][n] || cols[c][n] || grid[r/3*3+c/3][n] {
					return false
				}
				rows[r][n], cols[c][n], grid[r/3*3+c/3][n] = true, true, true
			}
		}
	}
	return true
}
