func convertToTitle(n int) string {
    var result []byte
    var a []int
    var t int
    var acc int
    
    for ; n > 26; {
        t = n%26
        if t != 0 {
            a = append(a, t)
            acc = 0
        } else {
            a = append(a, 26)
            acc = -1
        }
        n = n/26 + acc
    }
    if n != 0 {
        a = append(a, n)
    }
    
    for i := len(a)-1; i >= 0; i-- {
        var b byte
        b = byte(int(a[i])-1+int('A'))
        result = append(result, b)
    }
    
    return string(result)
}
