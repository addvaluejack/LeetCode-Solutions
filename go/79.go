func foo(board [][]byte, word string, x int, y int, index int, mark [][]int) bool {
    if len(word) == index {
        return true
    }
    // up
    if x > 0 && board[x-1][y] == byte(word[index]) && mark[x-1][y] == 0 {
        mark[x-1][y] = 1
        if foo(board, word, x-1, y, index+1, mark) == true {
            return true
        }
        mark[x-1][y] = 0
    }
    // down
    if x < len(board)-1 && board[x+1][y] == byte(word[index]) && mark[x+1][y] == 0 {
        mark[x+1][y] = 1
        if foo(board, word, x+1, y, index+1, mark) == true {
            return true
        }
        mark[x+1][y] = 0
    }
    // left
    if y > 0 && board[x][y-1] == byte(word[index]) && mark[x][y-1] == 0 {
        mark[x][y-1] = 1
        if foo(board, word, x, y-1, index+1, mark) == true {
            return true
        }
        mark[x][y-1] = 0
    }
    // right
    if y < len(board[0])-1 && board[x][y+1] == byte(word[index]) && mark[x][y+1] == 0 {
        mark[x][y+1] = 1
        if foo(board, word, x, y+1, index+1, mark) == true {
            return true
        }
        mark[x][y+1] = 0
    }
    return false
}

func exist(board [][]byte, word string) bool {
    if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == byte(word[0]) {
                mark := make([][]int, len(board))
                for i := 0; i < len(board); i++ {
                    mark[i] = make([]int, len(board[0]))
                }
                mark[i][j] = 1
                if foo(board, word, i, j, 1, mark) == true {
                    return true
                }
            }
        }
    }
    return false
}
