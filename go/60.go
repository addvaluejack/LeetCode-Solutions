func foo(n int) int {
    result := 1
    
    for i := 1; i <= n; i++ {
        result *= i
    }
    
    return result
}

func getPermutation(n int, k int) string {
    result := make([]byte, n)
    mark := make([]bool, n)
    
    k--
    for i := n; i > 0; i-- {
        tmp := k/foo(i-1)
        acc := 0
        for j := 0; j < n; j++ {
            if mark[j] == false && acc == tmp {
                result[n-i] = byte(int('1')+j)
                mark[j] = true
                break
            }
            if mark[j] == true {
                continue
            } else {
                acc++
            }
        }
        k = k%foo(i-1)
    }
    
    return string(result)
}
