func containsDuplicate(nums []int) bool {
    var m [1024][]int
    
    for i := 0; i < 1024; i++ {
        m[i] = make([]int, 0);
    }
    
    for i := 0; i < len(nums); i++ {
        var index int
        if nums[i] >= 0 {
            index = nums[i]%1024
        } else {
            index = (-nums[i])%1024
        }

        for j := 0; j < len(m[index]); j++ {
            if nums[i] == m[index][j] {
                return true
            }
        }
        m[index] = append(m[index], nums[i])
    }
    
    return false
}
