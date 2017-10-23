func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    if n == 0 {
        return
    }
    mark_row := make([]bool, m)
    mark_col := make([]bool, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                mark_row[i] = true
                mark_col[j] = true
            }
        }
    }
    
    for i := 0; i < m; i++ {
        if mark_row[i] == true {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }
    
    for j := 0; j < n; j++ {
        if mark_col[j] == true {
            for i := 0; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }
}
