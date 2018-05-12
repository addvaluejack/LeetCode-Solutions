func findLength(A []int, B []int) int {
    mark := make([][]int, len(A)+1)
    max := 0
    for i := 0; i <= len(A); i++ {
        mark[i] = make([]int, len(B)+1)
    }
    for i := 1; i <= len(A); i++ {
        for j := 1; j <= len(B); j++ {
            if A[i-1] == B[j-1] {
                mark[i][j] = mark[i-1][j-1]+1
            } else {
                mark[i][j] = 0
            }
            if mark[i][j] > max {
                max = mark[i][j]
            }
        }
    }
    
    return max
}
