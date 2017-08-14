func deepCopy (a [][]byte, b [][]byte) {
    for i := 0; i < 9; i++ {
        a[i] = make([]byte, 9)
        for j := 0; j < 9; j++ {
            a[i][j] = b[i][j]
        }
    }
}

func findExcatAnswer(board [][]byte, x int, y int) bool {
    var mark [9]int
    
    // check row
    for j := 0; j < 9; j++ {
        if board[x][j] == byte('.') {
            continue
        } else {
            if mark[int(board[x][j])-int('1')] == 0 {
                mark[int(board[x][j])-int('1')] = 1    
            }
        }
    }
    // check column
    for i := 0; i < 9; i++ {
        if board[i][y] == byte('.') {
            continue
        } else {
            if mark[int(board[i][y])-int('1')] == 0 {
                mark[int(board[i][y])-int('1')] = 1    
            }
        }
    }
    // check sub-box
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[x/3*3+i][y/3*3+j] == byte('.') {
                continue
            } else {
                if mark[int(board[x/3*3+i][y/3*3+j])-int('1')] == 0 {
                    mark[int(board[x/3*3+i][y/3*3+j])-int('1')] = 1    
                }
            }
        }
    }
    // if found, fill it. if not, do nothing
    acc := 0
    for i := 0; i < 9; i++ {
        if mark[i] == 1 {
            acc++
        }
    }
    if acc == 8 {
        for i := 0; i < 9; i++ {
            if mark[i] == 0 {
                board[x][y] = byte(i+int('1'))
            }
        }
        return true
    } else {
        return false
    }
}

func findPossibleAnswer(board [][]byte, tmp [][]byte, x int, y int){
    var mark [9]int
    
    // check row
    for j := 0; j < 9; j++ {
        if tmp[x][j] == byte('.') {
            continue
        } else {
            if mark[int(tmp[x][j])-int('1')] == 0 {
                mark[int(tmp[x][j])-int('1')] = 1    
            }
        }
    }
    // check column
    for i := 0; i < 9; i++ {
        if tmp[i][y] == byte('.') {
            continue
        } else {
            if mark[int(tmp[i][y])-int('1')] == 0 {
                mark[int(tmp[i][y])-int('1')] = 1    
            }
        }
    }
    // check sub-box
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if tmp[x/3*3+i][y/3*3+j] == byte('.') {
                continue
            } else {
                if mark[int(tmp[x/3*3+i][y/3*3+j])-int('1')] == 0 {
                    mark[int(tmp[x/3*3+i][y/3*3+j])-int('1')] = 1    
                }
            }
        }
    }
    // iteration
    for i := 0; i < 9; i++ {
        if mark[i] == 0 {
            tmp[x][y] = byte(i+int('1'))
            var m, n int
            flag := false
            for m = 0; m < 9; m++ {
                for n = 0; n < 9; n++ {
                    if tmp[m][n] == byte('.') {
                        flag = true
                        break
                    }
                }
                if flag {
                    break
                }
            }
            if m < 9 {
                clone := make([][]byte, 9)
                deepCopy(clone, tmp)
                findPossibleAnswer(board, clone, m, n)
            } else {
                deepCopy(board, tmp)
                board = tmp
                break
            }
        }
    }
}

func solveSudoku(board [][]byte)  {
    last_found_count := 0
    for ;; {
        current_found_count := 0
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                if board[i][j] == byte('.') {
                    if findExcatAnswer(board, i, j) {
                        current_found_count++
                    }
                } else {
                    current_found_count++
                }
            }
        }
        if current_found_count == last_found_count {
            break
        } else {
            last_found_count = current_found_count
        }
    }
    tmp := make([][]byte, 9)
    deepCopy(tmp, board)
    flag := false
    var m, n int
    for m = 0; m < 9; m++ {
        for n = 0; n < 9; n++ {
            if tmp[m][n] == byte('.') {
                flag = true
                break
            }
        }
        if flag {
            break
        }
    }
    if m < 9 {
        findPossibleAnswer(board, tmp, m, n);        
    }
}
