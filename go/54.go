func spiralOrder(matrix [][]int) []int {
    result := make([]int, 0)
    if len(matrix) == 0 {
        return result
    }
    top := 0
    bottom := len(matrix)-1
    left := 0
    right := len(matrix[0])-1
    
    for ; top < bottom && left < right; {
        for i := left; i < right; i++ {
            result = append(result, matrix[top][i])
        }
        for i := top; i < bottom; i++ {
            result = append(result, matrix[i][right])
        }
        for i := right; i > left; i-- {
            result = append(result, matrix[bottom][i])
        }
        for i := bottom; i > top; i-- {
            result = append(result, matrix[i][left])
        }
        top++
        bottom--
        left++
        right--
    }
    
    if top == bottom && left == right {
        result = append(result, matrix[top][left])
        return result
    }
    
    if top == bottom {
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
    }
    
    if left == right {
        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][left])
        }
        return result
    }
    
    return result
}
