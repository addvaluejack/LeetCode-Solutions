func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    
    result := 0
    mark := make([]bool, len(nums))
    mark[0] = true
    queue := make([]int, 0)
    queue = append(queue, 0)
    
    for ; len(queue) != 0; {
        result++
        new_queue := make([]int, 0)
        for i := 0; i < len(queue); i++ {
            for j := 1; j <= nums[queue[i]] && queue[i]+j < len(nums); j++ {
                if queue[i]+j == len(nums)-1 {
                    return result
                }
                if mark[queue[i]+j] == false {
                    new_queue = append(new_queue, queue[i]+j)
                    mark[queue[i]+j] = true
                }
            }
        }
        queue = new_queue
    }
    
    return -1
}
