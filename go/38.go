import (
    "bytes"
)

func countAndSay(n int) string {
    result := "1"
    
    for i := 1; i < n; i++ {
        var buffer bytes.Buffer
        
        for j := 0; j < len(result); j++ {
            count := 1
            for j+1 < len(result) && result[j+1] == result[j] {
                count++
                j++
            }
            buffer.WriteString(string(int('0')+count))
            buffer.WriteString(string(result[j]))
        }
        
        result = buffer.String()
    }
    
    return result
}
