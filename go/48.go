func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n/2; i++ {
        for j := 0; j < n-1-2*i; j++ {
            tmp := matrix[i][i+j]
            matrix[i][i+j] = matrix[n-1-i-j][i]
            matrix[n-1-i-j][i] = matrix[n-1-i][n-1-i-j]
            matrix[n-1-i][n-1-i-j] = matrix[i+j][n-1-i]
            matrix[i+j][n-1-i] = tmp
        }
    } 
}
