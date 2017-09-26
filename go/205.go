func isIsomorphic(s string, t string) bool {
    s_int := make([]int, len(s))
    t_int := make([]int, len(t))
    difference := make([]int, len(s))
    
    for i := 0; i < len(s); i++ {
        s_int[i] = int(s[i])
        t_int[i] = int(t[i])
        difference[i] = s_int[i] - t_int[i]
    }
    
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s); j++ {
            if s_int[i] == s_int[j] && difference[i] != difference[j] {
                return false
            }
            if t_int[i] == t_int[j] && difference[i] != difference[j] {
                return false
            }
        }
    }
    
    return true
}
