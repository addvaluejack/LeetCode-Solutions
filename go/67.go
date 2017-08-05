func addBinary(a string, b string) string {
    carry := 0
    var buffer bytes.Buffer
    bigger_len := len(b)
    if len(a) > len(b) {
        bigger_len = len(a)
    }
    
    for i := 0; i < bigger_len; i++ {
        elem_a := 0
        elem_b := 0
        
        if (len(a)-i-1 >= 0) && (a[len(a)-i-1] == '1') {
            elem_a = 1
        }
        if (len(b)-i-1 >= 0) && (b[len(b)-i-1] == '1') {
            elem_b = 1
        }
        
        if (elem_a+elem_b+carry)%2 == 1 {
            buffer.WriteString("1")
        } else {
            buffer.WriteString("0")
        }
        
        if (elem_a+elem_b+carry) >= 2 {
            carry = 1
        } else {
            carry = 0
        }
    }
    
    if carry == 1 {
        buffer.WriteString("1");
    }
    

    result := []rune(buffer.String())
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result);
}
