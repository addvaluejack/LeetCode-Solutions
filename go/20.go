func isValid(s string) bool {
    stack := make([]byte, 0)
    var tmp byte
    
    for i := 0; i < len(s); i++ {
        switch byte(s[i]) {
            case '(': {
                stack = append(stack, byte(s[i]))
            }
            case ')': {
                if len(stack) == 0 {
                    return false
                }
                stack, tmp = stack[:len(stack) - 1], stack[len(stack) - 1]
                if tmp != '(' {
                    return false
                }
            }
            case '{': {
                stack = append(stack, byte(s[i]))
            }
            case '}': {
                if len(stack) == 0 {
                    return false
                }
                stack, tmp = stack[:len(stack) - 1], stack[len(stack) - 1]
                if tmp != '{' {
                    return false
                }
            }
            case '[': {
                stack = append(stack, byte(s[i]))
            }
            case ']': {
                if len(stack) == 0 {
                    return false
                }
                stack, tmp = stack[:len(stack) - 1], stack[len(stack) - 1]
                if tmp != '[' {
                    return false
                }
            }
        }
    }
    
    if len(stack) != 0 {
        return false
    } else {
        return true
    }
}
