func findAnagrams(s string, p string) []int {
    result := make([]int, 0)
    p_mark := make([]int, 26)
    
    for i := 0; i < len(p); i++ {
        p_mark[int(p[i])-int('a')]++
    }
    
    for i := 0; i < len(s); i++ {
        t_mark := make([]int, 26)
        j := i
        for ; j < len(s) && j < i+len(p); j++ {
            t_mark[int(s[j])-int('a')]++
            if t_mark[int(s[j])-int('a')] > p_mark[int(s[j])-int('a')] {
                break
            }
        }
        if j == i+len(p) {
            result = append(result, i)
        }
    }
    
    return result
}
