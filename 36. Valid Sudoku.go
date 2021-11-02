// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
    row := make(map[int][]byte)
    col := make(map[int][]byte)
    subbox := make(map[int][]byte)
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            val := board[i][j]
            if val == '.' {
                continue
            }
            
            if inSlice(row[i], val) {
                return false
            }
            
            if inSlice(col[j], val) {
                return false
            }
            
            sidx := squareIndex(i, j)
            if inSlice(subbox[sidx], val) {
                return false
            }
            
            row[i] = append(row[i], val)
            col[j] = append(col[j], val)
            subbox[sidx] = append(subbox[sidx], val)
        }
    }
    return true
}

func inSlice(bs []byte, val byte) bool {
    if bs == nil {
        return false
    }
    for _, b := range bs {
        if b == val {
            return true
        }
    }
    return false
}

func squareIndex(i, j int) int {
    return (i / 3) * 10 + (j / 3)
}
