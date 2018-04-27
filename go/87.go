func check(s1 string, s2 string) bool {
    mark1 := make([]int, 26)
    mark2 := make([]int, 26)
    for i := 0; i < len(s1); i++ {
        mark1[int(s1[i])-int('a')]++
        mark2[int(s2[i])-int('a')]++
    }
    for i := 0; i < 26; i++ {
        if mark1[i] != mark2[i] {
            return false
        }
    }
    return true
}

func isScramble(s1 string, s2 string) bool {
    if check(s1, s2) == false {
        return false
    }
    if len(s1) == 1 && s1 != s2 {
        return false
    }
    if s1 == s2 {
        return true
    }
    for i := 1; i < len(s1); i++ {
        if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
            return true
        }
        if isScramble(s1[:i], s2[len(s1)-i:]) && isScramble(s1[i:], s2[:len(s1)-i]) {
            return true
        }
    }
    return false
}
