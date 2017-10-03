func matchMark(mark [26]int, marks[][26]int) int {
    for i := 0; i < len(marks); i++ {
        j := 0
        for ; j < 26; j++ {
            if mark[j] != marks[i][j] {
                break
            }
        }
        if j == 26 {
            return i
        }
    }
    
    return -1
}

func groupAnagrams(strs []string) [][]string {
    var results [][]string
    var marks [][26]int
    
    for i := 0; i < len(strs); i++ {
        var mark [26]int
        
        for j := 0; j < len(strs[i]); j++ {
            mark[int(strs[i][j])-int('a')] += 1
        }
        
        marks_index := matchMark(mark, marks)
        
        if marks_index == -1 {
            marks = append(marks, mark)
            result := make([]string, 1)
            result[0] = strs[i]
            results = append(results, result)
        } else {
            results[marks_index] = append(results[marks_index], strs[i])
        }
    }
    
    return results
}
