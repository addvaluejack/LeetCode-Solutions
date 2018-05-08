func min(valueA int, valueB int, valueC int) int {
    if valueA <= valueB && valueA <= valueC {
        return valueA
    }
    if valueB <= valueC && valueB <= valueA {
        return valueB
    }
    return valueC
}

func minDistance(word1 string, word2 string) int {
    if len(word1) == 0 {
        return len(word2)
    }
    if len(word2) == 0 {
        return len(word1)
    }
    mark := make([][]int, len(word1)+1)
    for i := 0; i <= len(word1); i++ {
        mark[i] = make([]int, len(word2)+1)
        mark[i][0] = i
    }
    for j := 0; j <= len(word2); j++ {
        mark[0][j] = j
    }
    for i := 1; i <= len(word1); i++ {
        for j := 1; j <= len(word2); j++ {
            if word1[i-1] == word2[j-1] {
                mark[i][j] = min(mark[i-1][j-1], mark[i-1][j]+1, mark[i][j-1]+1)
            } else {
                mark[i][j] = min(mark[i-1][j-1]+1, mark[i-1][j]+1, mark[i][j-1]+1)
            }
        }
    }
    return mark[len(word1)][len(word2)]
}
