func numJewelsInStones(J string, S string) int {
    mark := make([]bool, 256)
    result := 0
    
    for i := 0; i < len(J); i++ {
        mark[int(J[i])] = true
    }
    
    for j := 0; j < len(S); j++ {
        if mark[int(S[j])] == true {
            result++
        }
    }
    
    return result
}
