func generateParenthesis(n int) []string {
    result := make([]string, 0)
    current := make([]byte, 0)
    
    foo(n, 0, current, &result)
    
    return result
}

func foo(n int, left int, current []byte, result_ptr *[]string) {
    if n == 0 && left == 0 {
        *result_ptr = append(*result_ptr, string(current))
        return
    }
    if n == 0 && left > 0 {
        tmp := append(current, ')')
        foo(n, left - 1, tmp, result_ptr)
    }
    if n > 0 {
        tmp1 := append(current, '(')
        foo(n - 1, left + 1, tmp1, result_ptr)
        if left > 0 {
            tmp2 := append(current, ')')
            foo(n, left - 1, tmp2, result_ptr)
        }
    }
    return
}
