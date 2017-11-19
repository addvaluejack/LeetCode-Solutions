func reverseVowels(s string) string {
    vowel_mark := make([]int, 0)
    result := []byte(s)
    
    for i := 0; i < len(s); i++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
            vowel_mark = append(vowel_mark, i)
        }
    }
    
    for i := 0; i < len(vowel_mark)/2; i++ {
        tmp := result[vowel_mark[i]]
        result[vowel_mark[i]] = result[vowel_mark[len(vowel_mark)-1-i]]
        result[vowel_mark[len(vowel_mark)-1-i]] = tmp
    }
    
    return string(result)
}
