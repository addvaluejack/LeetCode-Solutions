func search(nums []int, target int) bool {  
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] < nums[i-1] && (target > nums[0] || target > nums[i-1]) {
            break
        }
        if nums[i] == target {
            return true
        }
    }
    
    return false
}
