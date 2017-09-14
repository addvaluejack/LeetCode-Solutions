func rotate(nums []int, k int) {
    k = k%len(nums)
    
    result := nums
    result = append(result[len(nums)-k:], result[:len(nums)-k]...)
    
    for i := 0; i < len(nums); i++ {
        nums[i] = result[i]
    }
}
