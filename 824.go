import("strings")

func processVovwel(s string) string {
    return s+"ma"
}

func processConsonant(s string) string {
    result := make([]byte, len(s))
    for i := 0; i < len(s)-1; i++ {
        result[i] = s[i+1]
    }
    result[len(s)-1] = s[0]
    return string(result)
}

func checkVovwel(b byte) bool {
    vovwels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for j := 0; j < 10; j++ {
        if b == vovwels[j] {
            return true
        }
    }
    return false
}

func toGoatLatin(S string) string {
    ss := strings.Split(S, " ")
    for i := 0; i < len(ss); i++ {
        flag := false
        if checkVovwel(ss[i][0]) {
            flag = true
        }
        if flag == false {
            ss[i] = processConsonant(ss[i])
        }
        ss[i] = processVovwel(ss[i])
        for j := 0; j <= i; j++ {
            ss[i] = ss[i]+"a"
        }
    }
    result := ""
    for i := 0; i < len(ss); i++ {
        result = result+" "+ss[i]
    }
    return result[1:]
}
