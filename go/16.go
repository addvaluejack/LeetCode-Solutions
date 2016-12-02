import(
    "sort"
    "math"
)

func threeSumClosest(nums []int, target int) int {
    var result int
    flag := false
    
    sort.Ints(nums)
    i := 0
    j := 1
    k := 2
    for ; i < len(nums); {
        for ; j < len(nums); {
            for ; k < len(nums); {
                if flag == false || math.Abs(float64(target) - float64(nums[i] + nums[j] + nums[k])) < math.Abs(float64(target) - float64(result)) {
                    result = nums[i] + nums[j] + nums[k]
                    flag = true
                }
                num_k := nums[k]
                for k < len(nums) && num_k == nums[k] {
                    k++
                }
            }
            num_j := nums[j]
            for j < len(nums) && num_j == nums[j] {
                j++
            }
	    k = j + 1
        }
        num_i := nums[i]
        for i < len(nums) && num_i == nums[i] {
            i++
        }
	j = i + 1
	k = j + 1
    }
 
    return result
}
