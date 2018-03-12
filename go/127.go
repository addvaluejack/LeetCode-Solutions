func isValid(A string, B string) bool {
    diff := 0
    for i := 0; i < len(A); i++ {
        if A[i] != B[i] {
            diff++
        }
    }
    
    return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    transformationCount := 1
    usedWordMark := make([]bool, len(wordList))
    currentWords := make([]string, 0)
    currentWords = append(currentWords, beginWord)
    
    for ; ; {
        newCurrentWords := make([]string, 0)
        for i := 0; i < len(currentWords); i++ {
            if currentWords[i] == endWord {
                return transformationCount
            }
            for j := 0; j < len(wordList); j++ {
                if usedWordMark[j] == false && isValid(currentWords[i], wordList[j]) {
                    usedWordMark[j] = true
                    newCurrentWords = append(newCurrentWords, wordList[j])
                }
            }
        }
        if len(newCurrentWords) == 0 {
            break
        }
        currentWords = newCurrentWords
        transformationCount++
    }
    
    return 0
}
