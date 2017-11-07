func isMatch(s string, p string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    
    for i := 0; i < len(p); i++ {
        if p[i] != '*' {
            for j := len(s)-1; j>= 0; j-- {
                dp[j+1] = dp[j] && (p[i] == s[j] || p[i] == '?')
            }
        } else {
            for j := 1; j < len(s)+1; j++ {
                dp[j] = dp[j-1] || dp[j]
            }
        }
        dp[0] = dp[0] && p[i] == '*'
    }
    
    return dp[len(s)]
}
