func isAlphaChar(a byte) bool {
    if int(a) >= int('a') && int(a) <= int('z') {
        return true
    }
    if int(a) >= int('A') && int(a) <= int('Z') {
        return true
    }
    if int(a) >= int('0') && int(a) <= int('9') {
        return true
    }
    return false
}

func lowCase(a byte) byte {
    if int(a) >= int('A') && int(a) <= int('Z') {
        return byte(int(a)+int('a')-int('A'))
    } else {
        return a
    }
}

func isPalindrome(s string) bool {
    i := 0
    j := len(s)-1
    for ; i < j; {
        for ; i < len(s) && !isAlphaChar(s[i]); i++ {
        }
        if i > j || i == len(s) {
            return true
        }
        for ; j >= 0 && !isAlphaChar(s[j]); j-- {
        }
        if lowCase(s[i]) != lowCase(s[j]) {
            return false
        }
        i++
        j--
    }
    
    return true
}
