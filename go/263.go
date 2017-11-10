func isUgly(num int) bool {
    if num == 0 {
        return false
    }
    
    for i := 6; i > 1; i-- {
        for ; num%i == 0; {
            num = num/i
        }
    }
    
    if num == 1 {
        return true
    } else {
        return false
    }
}
