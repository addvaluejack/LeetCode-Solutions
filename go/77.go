func foo(n int, k int, used int, previous []int, index int, result_ptr *[][]int) {
    if index == k {
        *result_ptr = append(*result_ptr, previous)
        return
    }
    for i := used; i <= n; i++ {
        current := make([]int, k)
        for j := 0; j < k; j++ {
            current[j] = previous[j]
        }
        current[index] = i
        foo(n, k, i+1, current, index+1, result_ptr)
    }
}

func combine(n int, k int) [][]int {
    result := make([][]int, 0)
    
    if n == 0 || k == 0 || k > n {
        return result
    }
    
    for i := 1; i <= n; i++ {
        current := make([]int, k)
        current[0] = i
        foo(n, k, i+1, current, 1, &result)
    }
    
    return result
}
