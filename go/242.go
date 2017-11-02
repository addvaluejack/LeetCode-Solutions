func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    s_mark := make([]int, 26)
    t_mark := make([]int, 26)
    
    for i := 0; i < len(s); i++ {
        s_mark[int(s[i])-int('a')]++
        t_mark[int(t[i])-int('a')]++
    }
    
    for i := 0; i < 26; i++ {
        if s_mark[i] != t_mark[i] {
            return false
        }
    }
    
    return true
}
