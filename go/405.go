func toHex(num int) string {
    reference := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
    result := make([]byte, 0)
    
    if num == 0 {
        return "0"
    } else if num > 0 {
        for ; num != 0; {
            tmp := num%16
            result = append(result, reference[tmp])
            num = num/16
        }
        for i := 0; i < len(result)/2; i++ {
            result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
        }
        return string(result)
    } else {
        num = num+1
        for ; num != 0; {
            tmp := num%16
            result = append(result, reference[15+tmp])
            num = num/16
        }
        for i := len(result); i < 8; i++ {
            result = append(result, 'f')
        }
        for i := 0; i < len(result)/2; i++ {
            result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
        }
        return string(result)
    }
}
