func addStrings(num1 string, num2 string) string {
    result := make([]byte, 0)
    carry := 0
    
    for i := 1; i <= len(num1) || i <= len(num2); i++ {
        ts := make([]byte, 1)
        t := 0
        if  i <= len(num1) && i <= len(num2) {
            t = int(num1[len(num1)-i])+int(num2[len(num2)-i])-int('0')*2+carry
        } else if i > len(num1) {
            t = int(num2[len(num2)-i])-int('0')+carry

        } else {
            t = int(num1[len(num1)-i])-int('0')+carry
            carry = 0
        }
        carry = t/10
        ts[0] = byte(t%10+int('0'))
        result = append(ts, result...)
    }
    
    if carry == 1 {
        ts := make([]byte, 1)
        ts[0] = '1'
        result = append(ts, result...)
    }
    
    return string(result)
}
