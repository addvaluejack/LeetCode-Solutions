func generateMatrix(n int) [][]int {
    result := make([][]int, n)
    
    for i := 0; i < n; i++ {
        result[i] = make([]int, n)
    }
    
    top := 0
    bottom := n-1
    left := 0
    right := n-1
    count := 1
    
    for ; left < right; {
        for i := left; i < right; i++ {
            result[top][i] = count
            count++
        }
        for i := top; i < bottom; i++ {
            result[i][right] = count
            count++
        }
        for i := right; i > left; i-- {
            result[bottom][i] = count
            count++
        }
        for i := bottom; i > top; i-- {
            result[i][left] = count
            count++
        }
        left++
        right--
        top++
        bottom--
    }
    
    if left == right {
        result[left][top] = count
    }
    
    return result
}
