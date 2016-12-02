import "sort"

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)
    
    if len(nums) < 3 {
        return result
    }
    
    sort.Ints(nums)
    i := 0
    j := 1
    k := 2
    for ; i < len(nums); {
        for ; j < len(nums); {
            for ; k < len(nums); {
                if nums[i] + nums[j] + nums[k] == 0 {
                    tmp := []int{nums[i], nums[j], nums[k]}
                    result = append(result, tmp)
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
    
    return result;
}
