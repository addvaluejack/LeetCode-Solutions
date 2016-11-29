func romanToInt(s string) int {
    intergers := [7]int{1, 5, 10, 50, 100, 500, 1000}
    romans := [7]byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
    
    result := 0
    
    for i := 6; i > -1; i-- {
        for j := 0; j < len(s); j++ {
            if(s[j] == romans[i]) {
                flag := false
                for k := i + 1; k < 7; k++ {
                    for l := j + 1; l < len(s); l++ {
                        if(s[l] == romans[k]) {
                            flag = true;
                            break;
                        }
                    }
                    if(flag == true) {
                        break;
                    }
                }
                if(flag == true) {
                    result -= intergers[i];
                } else {
                    result += intergers[i];
                }
            }
        }
    }
    
    return result;
}
