func trailingZeroes(n int) int {
    result := 0
    divide := 5
    
    for ;; {
        tmp := n/divide
        if tmp == 0 {
            break
        } else {
            result += tmp
            divide *= 5
        }
    }
    
    return result
}
