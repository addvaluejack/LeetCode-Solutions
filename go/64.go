func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    path := make([][]int, m)
    
    for i := 0; i < m; i++ {
        path[i] = make([]int, n)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                path[i][j] = grid[i][j]
            }
            if i > 0 && j == 0 {
                path[i][j] = grid[i][j]+path[i-1][j]
            }
            if i == 0 && j > 0 {
                path[i][j] = grid[i][j]+path[i][j-1]
            }
            if i > 0 && j > 0 {
                if path [i-1][j] < path[i][j-1] {
                    path[i][j] = grid[i][j]+path[i-1][j]
                } else {
                    path[i][j] = grid[i][j]+path[i][j-1]
                }
            }
        }
    }
    
    return path[m-1][n-1]
}
