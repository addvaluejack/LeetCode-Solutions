func firstUniqChar(s string) int {
    count := make([]int, 26)
    index := make([]int, 26)
    
    for i := 0; i < 26; i++ {
        index[i] = -1
    }
    
    for i := 0; i < len(s); i++ {
        if count[int(s[i])-int('a')] == 0 && index[int(s[i])-int('a')] == -1 {
            index[int(s[i])-int('a')] = i
        }
        count[int(s[i])-int('a')]++
    }
    
    result := -1
    for i := 0; i < 26; i++ {
        if count[i] == 1 {
            if result == -1 {
                result = index[i]
            } else if index[i] < result {
                result = index[i]
            }
        }
    }
    
    return result
}
