func sortColors(nums []int)  {
    if len(nums) == 0 {
        return 
    }
    
    i, j := 0, 0
    
    for k := 0; k < len(nums); k++ {
        tmp := nums[k]
        nums[k] = 2
        if tmp < 2 {
            nums[j] = 1
            j++
        }
        if tmp == 0 {
            nums[i] = 0
            i++
        }
    }
}
