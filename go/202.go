func isHappy(n int) bool {
    fast_n := n
    slow_n := n
    for ;; {
        slow_n = getNewN(slow_n)
        fast_n = getNewN(fast_n)
        fast_n = getNewN(fast_n)
        if slow_n == fast_n {
            break
        }
    }
    if slow_n == 1 {
        return true
    }
    
    return false
}

func getNewN(n int) int {
    new_n := 0
    for ; n >= 10; {
        new_n += pow((n%10), 2)
        n /= 10
    }
    new_n += pow(n, 2)
    return new_n
}

func pow(x int, y int) int {
    result := 1
    for i := 0; i < y; i++ {
        result *= x
    }
    
    return result
}
