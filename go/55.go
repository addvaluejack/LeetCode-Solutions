func canJump(nums []int) bool {
    current := 0
    fartherest := current+nums[current]
      
    for ; current <= fartherest; current++ {
        if current == len(nums)-1 {
            return true
        }
        if current+nums[current] > fartherest {
            fartherest = current+nums[current]
        }
    }
    
    return false
}
