func removeDuplicates(nums []int) int {
    i := 0
    j := 0
    
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return 1
    }
    
    for {
        for j < len(nums) && nums[j] == nums[i] {
            j++
        }
        if j == len(nums) {
            break
        } else {
            i++
            nums[i] = nums[j]
        }
    }
    
    return i+1
}
