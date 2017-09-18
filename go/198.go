func rob(nums []int) int {
    result := make([]int, len(nums))
    
    if len(nums) == 0 {
        return 0
    }
    
    for i :=0; i < len(nums); i++ {
        tmp := 0
        if i-3 >= 0 {
            tmp = result[i-3]
            if result[i-2] > result[i-3] {
                tmp = result[i-2]
            }
        } else if i-2 >= 0{
            tmp = result[i-2]
        }
        result[i] = tmp + nums[i]
    }
    
    if len(nums) >= 2 && result[len(nums)-2] > result[len(nums)-1] {
        return result[len(nums)-2]
    }
    
    return result[len(nums)-1]
}
