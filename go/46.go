func foo(result_ptr *[][]int, nums []int, mark []int, index int) {
    if index == len(nums)-1 {
        permutation := make([]int, len(nums))
        for i := 0; i < len(nums); i++ {
            permutation[i] = nums[mark[i]];
        }
        (*result_ptr) = append((*result_ptr), permutation);
        return
    }
    for i := 0; i < len(nums); i++ {
        if mark[i] == -1 {
            new_mark := make([]int, len(nums))
            for j := 0; j < len(nums); j++ {
                new_mark[j] = mark[j]
            }
            new_mark[i] = index+1
            foo(result_ptr, nums, new_mark, index+1)
        }
    }
}

func permute(nums []int) [][]int {
    var result [][]int
    
    if len(nums) == 0 {
        var tmp []int
        return append(result, tmp)
    }
    
    mark := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        mark[i] = -1
    }
    
    for i := 0; i < len(nums); i++ {
        mark[i] = 0
        foo(&result, nums, mark, 0)
        mark[i] = -1
    }
    
    return result
}
