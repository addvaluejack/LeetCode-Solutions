func pow(x int, y int) int {
    result := 1
    
    for i := 0; i < y; i++ {
        result = result*x
    }
    
    return result
}

func grayCode(n int) []int {
    result := make([]int, 0)

    result = append(result, 0)
    if n == 0 {
        return result
    }
    
    result = append(result, 1)
    
    if n == 1 {
        return result
    }
    
    for i := 2; i <= n; i++ {
        for j := pow(2, i-1); j < pow(2, i); j++ {
            tmp := pow(2, i-1)+result[pow(2, i)-1-j]
            result = append(result, tmp)
        }
    }
    
    return result
}
