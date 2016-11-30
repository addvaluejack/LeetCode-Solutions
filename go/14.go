import(
    "sort"
    "bytes"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func longestCommonPrefix(strs []string) string {
    var result bytes.Buffer
    
    if len(strs) == 0 {
        return ""
    }
    
    sort.Sort(ByLength(strs))
    
    for i:= 0; i < len(strs[0]); i++ {
        flag := false
        for j:= 0; j < len(strs); j++ {
            if strs[0][i] != strs[j][i] {
                flag = true
                break
            }
        }
        if flag == false {
            result.WriteByte(strs[0][i])
        } else {
            break
        }
    }
    
    return result.String()
}
