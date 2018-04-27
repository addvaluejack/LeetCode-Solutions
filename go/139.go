func wordBreak(s string, wordDict []string) bool {
    mark := make([]bool, len(s)+1)
    mark[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < len(wordDict); j++ {
            if i-len(wordDict[j]) >= 0 {
                if s[i-len(wordDict[j]):i] == wordDict[j] && mark[i-len(wordDict[j])] == true {
                    mark[i] = true
                }
            }
        }
    }
    return mark[len(s)]
}
