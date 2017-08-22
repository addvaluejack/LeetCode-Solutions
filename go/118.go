func generate(numRows int) [][]int {
    var result [][]int
    
    if numRows == 0 {
        return result
    }
    
    first_level := []int{1}
    result = append(result, first_level)
    
    for i := 2; i <= numRows; i++ {
        current_level := make([]int, i)
        for j := 0; j < i; j++ {
            if j == 0 {
                current_level[j] = 1
            } else if j == i-1 {
                current_level[j] = 1
            } else {
                current_level[j] = result[i-2][j-1]+result[i-2][j]
            }
        }
        result = append(result, current_level)
    }
    
    return result
}
