func calculate(s string) int {
    result := 0
    minus_sign := false
    minus_layers := make([]bool, 1)
    
    for i := 0; i < len(s); i++ {
        if int(s[i]) >= int('0') && int(s[i]) <= int('9') {
            tmp := int(s[i]) - int('0')
            for ; i+1 < len(s) && int(s[i+1]) >= int('0') && int(s[i+1]) <= int('9'); {
                i++
                tmp = tmp*10 + int(s[i]) - int('0')
            }
            if minus_sign == minus_layers[len(minus_layers)-1] {
                result += tmp
            } else {
                result -= tmp
            }
        }
        if s[i] == '-' {
            minus_sign = true
        }
        if s[i] == '+' {
            minus_sign = false
        }
        if s[i] == '(' {
            if minus_sign {
                minus_layers = append(minus_layers, !minus_layers[len(minus_layers)-1])
                minus_sign = false
            } else {
                minus_layers = append(minus_layers, minus_layers[len(minus_layers)-1])
            }
        }
        if s[i] == ')' {
            minus_layers = minus_layers[:len(minus_layers)-1]
        }
    }
    
    return result
}
