// https://leetcode.com/problems/rotting-oranges/
func orangesRotting(grid [][]int) int {
    rottens := make([][]int, 0)
    
    freshCount := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            c := grid[i][j]
            if c == 2 {
                rottens = append(rottens, []int{i,j})
            } else if c == 1 {
                freshCount++
            }
        }
    }
    if freshCount == 0 {
        return 0
    }
    
    cnt := 0
    dirs := [][]int { {0,1}, {0,-1}, {1,0}, {-1,0}}
    for len(rottens) > 0 {
        cnt++
        s := len(rottens)
        for i:=0; i < s; i++ {
            r := rottens[0]
            rottens = rottens[1:]
            for _, d := range dirs {
                if rotte(grid, (r[0] + d[0]), (r[1] + d[1])) {
                    freshCount--
                    rottens = append(rottens, []int{r[0] + d[0], r[1] + d[1]})
                }
            }
        }
        if freshCount == 0 {
            return cnt
        }
    }

    // remain fresh orange, return -1
    if freshCount > 0 {
        return -1
    }
    return cnt-1
}

func rotte(grid [][]int, i, j int) bool {
    if validBoundry(grid, i, j) && grid[i][j] == 1 {
        grid[i][j] = 2
        return true
    }
    return false
}

func validBoundry(grid [][]int, i, j int) bool {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
        return false
    }
    return true
}
