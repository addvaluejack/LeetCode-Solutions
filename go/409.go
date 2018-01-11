func longestPalindrome(s string) int {
    count := make([]int, 256)
    result := 0
    flag := false
    
    for i := 0; i < len(s); i++ {
        count[int(s[i])]++
    }
    
    for i := 0; i < 256; i++ {
        if count[i]%2 == 1 {
            result += count[i]-1
            flag = true
        } else {
            result += count[i]
        }
    }
    
    if flag == true {
        return result+1
    } else {
        return result
    }
}
