func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obstacleGrid[0])
    path := make([][]int, m)
    
    for i := 0; i < m; i++ {
        path[i] = make([]int, n)
    }
    
    path[0][0] = 1
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                path[i][j] = 0
                continue
            }
            if i > 0 {
                path[i][j] += path[i-1][j]
            }
            if j > 0 {
                path[i][j] += path[i][j-1]
            }
        }
    }
    
    return path[m-1][n-1]
}
