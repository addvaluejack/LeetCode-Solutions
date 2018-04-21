func foo(board_ptr *[][]byte, x int, y int) {
    if (*board_ptr)[x][y] == '-' || (*board_ptr)[x][y] == 'X' {
        return
    }
    (*board_ptr)[x][y] = '-'
    if x > 0 && (*board_ptr)[x-1][y] == 'O' {
        foo(board_ptr, x-1, y)
    }
    if x < len(*board_ptr)-1 && (*board_ptr)[x+1][y] == 'O' {
        foo(board_ptr, x+1, y)
    }
    if y > 0 && (*board_ptr)[x][y-1] == 'O' {
        foo(board_ptr, x, y-1)
    }
    if y < len((*board_ptr)[0])-1 && (*board_ptr)[x][y+1] == 'O' {
        foo(board_ptr, x, y+1)
    }
}

func solve(board [][]byte)  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
                foo(&board, i, j)
            }
        }
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            if board[i][j] == '-' {
                board[i][j] = 'O'
            }
        }
    }
}
