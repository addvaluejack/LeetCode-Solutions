func removeElement(nums []int, val int) int {
    i, j := 0, len(nums)
    
    for i < j {
        if nums[i] == val {
            for k := i; k < j - 1; k++ {
                nums[k] = nums[k + 1]
            }
            j--
        } else {
            i++
        }
    }
    
    return j
}
