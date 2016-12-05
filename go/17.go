import "bytes"

var letters = [][]string{
    {" "},
    {""},
    {"a", "b", "c"},
    {"d", "e", "f"},
    {"g", "h", "i"},
    {"j", "k", "l"},
    {"m", "n", "o"},
    {"p", "q", "r", "s"},
    {"t", "u", "v"},
    {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
    result := make([]string, 0)
    var current bytes.Buffer    
    
    if(len(digits) == 0) {
        var nothing []string
        return nothing
    }

    foo(digits, 0, current, &result)
    
    return result
}

func foo(digits string, index int, current bytes.Buffer, result_pointer *[]string) {
    if index == len(digits) {
        *result_pointer = append(*result_pointer, current.String())
        return
    }
    for i := 0; i < len(letters[int(digits[index] - '0')]); i++ {
        var new_current bytes.Buffer
        var new_index int
        new_current = current
        new_current.WriteString(letters[int(digits[index] - '0')][i])
        new_index = index + 1
        foo(digits, new_index, new_current, result_pointer)
    }
}
