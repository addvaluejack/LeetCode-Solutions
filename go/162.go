func findPeakElement(nums []int) int {
    result := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            result++
        } else {
            break
        }
    }
    return result
}
