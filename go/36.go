func isValidSudoku(board [][]byte) bool {
    // check row
    for i := 0; i < 9; i++ {
        var mark [9]int
        for j := 0; j < 9; j++ {
            if board[i][j] == byte('.') {
                continue
            }
            mark[int(board[i][j])-int('0')-1]++
            if mark[int(board[i][j])-int('0')-1] > 1 {
                return false
            }
        }
    }
    // check column
    for i := 0; i < 9; i++ {
        var mark [9]int
        for j := 0; j < 9; j++ {
            if board[j][i] == byte('.') {
                continue
            }
            mark[int(board[j][i])-int('0')-1]++
            if mark[int(board[j][i])-int('0')-1] > 1 {
                return false
            }
        }
    }
    // check sub-box
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            var mark [9]int
            for k := 0; k < 3; k++ {
                for l := 0; l < 3; l++ {
                    if board[i*3+k][j*3+l] == byte('.') {
                        continue
                    }
                    mark[int(board[i*3+k][j*3+l])-int('0')-1]++
                    if mark[int(board[i*3+k][j*3+l])-int('0')-1] > 1 {
                        return false
                    }
                }
            }
        }
    }
    
    return true
}
