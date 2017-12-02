func removeDuplicates(nums []int) int {
    result := make([]int, 0)
    twice := false
    
    for i := 0; i < len(nums); i++{
        if i == 0 {
            result = append(result, nums[0])
        } else {
            if nums[i] == result[len(result)-1] {
                if twice == true {
                    continue
                } else {
                    result = append(result, nums[i])
                    twice = true
                }
            } else {
                result = append(result, nums[i])
                twice = false
            }
        }
    }
    
    for i := 0; i < len(result); i++ {
        nums[i] = result[i]
    }
    return len(result)
}
