import (
    "strconv"
)

func restoreIpAddresses(s string) []string {
    result := make([]string, 0)
    sets := make([]int, 0)
    
    foo(s, 0, sets, &result)
    
    return result
}

func foo(s string, index int, sets []int, result_ptr *[]string) {
    if len(sets) > 4 {
        return
    }
    if index == len(s) && len(sets) == 4 {
        tmp := strconv.Itoa(sets[0])+"."+strconv.Itoa(sets[1])+"."+strconv.Itoa(sets[2])+"."+strconv.Itoa(sets[3])
        *result_ptr = append(*result_ptr, tmp)
        return
    }
    if index+2 < len(s) && s[index] == '2' && s[index+1] == '5' && int(s[index+2])-int('0') <= 5 {
        t_sets := sets
        tmp := 250+(int(s[index+2])-int('0'))
        t_sets = append(t_sets, tmp)
        foo(s, index+3, t_sets, result_ptr)
    }
    if index+2 < len(s) && s[index] == '2' && int(s[index+1])-int('0') <= 4 {
        t_sets := sets
        tmp := 200+(int(s[index+1])-int('0'))*10+(int(s[index+2])-int('0'))
        t_sets = append(t_sets, tmp)
        foo(s, index+3, t_sets, result_ptr)        
    }
    if index+2 < len(s) && s[index] == '1' {
        t_sets := sets
        tmp := 100+(int(s[index+1])-int('0'))*10+(int(s[index+2])-int('0'))
        t_sets = append(t_sets, tmp)
        foo(s, index+3, t_sets, result_ptr)    
    }
    if index+1 < len(s) && s[index] != '0' {
        t_sets := sets
        tmp := (int(s[index])-int('0'))*10+(int(s[index+1])-int('0'))
        t_sets = append(t_sets, tmp)
        foo(s, index+2, t_sets, result_ptr)
    }
    if index < len(s) {
        t_sets := sets
        tmp := int(s[index])-int('0')
        t_sets = append(t_sets, tmp)
        foo(s, index+1, t_sets, result_ptr)
    }
}

