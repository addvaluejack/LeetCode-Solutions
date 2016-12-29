func strStr(haystack string, needle string) int {
    result := -1
    
    if len(haystack) == 0 && len(needle) == 0 {
        return 0
    }
    
    for i := 0; i < len(haystack); i++ {
        j := 0
        for ; j < len(needle); j++ {
            if i+j+1 > len(haystack) || haystack[i+j] != needle[j] {
                break;
            }
        }
        if j == len(needle) {
            return i
        }
    }
    
    return result;
}
