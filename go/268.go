func missingNumber(nums []int) int {
    result := (0+len(nums))*(len(nums)+1)/2

    for i := 0; i < len(nums); i++ {
        result -= nums[i]
    }
    
    return result
}
