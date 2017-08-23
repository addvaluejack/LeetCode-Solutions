func getRow(rowIndex int) []int {
    var result [2][]int
    result[0] = make([]int, rowIndex+1)
    result[1] = make([]int, rowIndex+1)
    
    result[0][0] = 1
    
    for i := 1; i <= rowIndex; i++ {
        if i%2 == 1 {
            for j := 0; j < i+1; j++ {
                if j == 0 {
                    result[1][j] = 1
                } else if j == i {
                    result[1][j] = 1
                } else {
                    result[1][j] = result[0][j-1]+result[0][j]
                }
            }
        } else {
            for j := 0; j < i+1; j++ {
                if j == 0 {
                    result[0][j] = 1
                } else if j == i {
                    result[0][j] = 1
                } else {
                    result[0][j] = result[1][j-1]+result[1][j]
                }
            }
        }
    }
    
    if rowIndex%2 == 1 {
        return result[1]
    } else {
        return result[0]
    }
}
