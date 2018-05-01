import "strings"

func calculateKey(word string) int {
    key := 0
    for i := len(word)-1; i >= 0; i-- {
        key *= 26
        key += int(word[i])-int('a')
    }
    return key
}

func wordPattern(pattern string, str string) bool {
    keys := make([]int, len(pattern))
    words := strings.Split(str, " ")
    if len(pattern) != len(words) {
        return false
    }
    result := true
    for i := 0; i < len(pattern); i++ {
        keys[i] = calculateKey(words[i])
    }
    for i := 0; i < len(pattern); i++ {
        for j := i+1; j < len(pattern); j++ {
            if pattern[i] == pattern[j] && keys[i] != keys[j] {
                result = false
                break
            } else if pattern[i] != pattern[j] && keys[i] == keys[j] {
                result = false
                break
            }
        }
        if result == false {
            break
        }
    }
    return result
}
