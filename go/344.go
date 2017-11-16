func reverseString(s string) string {
    result := make([]byte, len(s))
    
    for i := 0; i < len(s); i++ {
        result[i] = s[len(s)-1-i]
    }
    
    return string(result)
}
