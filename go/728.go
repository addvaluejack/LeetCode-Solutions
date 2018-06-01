func foo(original int, current int) bool {
    if current == 0 {
        return true
    }
    
    t := current%10
    if t == 0 || original%t != 0 {
        return false    
    } else {
        return foo(original, current/10)
    }
}

func selfDividingNumbers(left int, right int) []int {
    result := make([]int, 0)
    for i := left; i <= right; i++ {
        if foo(i, i) == true {
            result = append(result, i)
        }
    }
    
    return result
}
