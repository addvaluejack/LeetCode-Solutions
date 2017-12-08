func foo(s string, index int, result_ptr *int) {
    if index == len(s) {
        (*result_ptr)++
        return
    }
    if s[index] == '0' {
        return
    }
    if s[index] == '1' && index+2 <= len(s) {
        foo(s, index+2, result_ptr)
    }
    if s[index] == '2' && index+2 <= len(s) && int(s[index+1])-int('0') <= 6 {
        foo(s, index+2, result_ptr)
    }
    foo(s, index+1, result_ptr)
}

func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    result := 0
    foo(s, 0, &result)
    
    return result
}
