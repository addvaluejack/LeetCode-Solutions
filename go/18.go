import "sort"

func fourSum(nums []int, target int) [][]int {
    result := make([][]int, 0)
    
    if len(nums) < 4 {
        return result
    }
    
    sort.Ints(nums)
    i := 0
    j := 1
    k := 2
    l := 3
    for ; i < len(nums); {
        for ; j < len(nums); {
            for ; k < len(nums); {
                for ; l < len(nums); {
                    if nums[i] + nums[j] + nums[k] + nums[l] == target {
                        tmp := []int{nums[i], nums[j], nums[k], nums[l]}
                        result = append(result, tmp)
                    }
                    num_l := nums[l]
                    for l < len(nums) && num_l == nums[l] {
                        l++
                    }
                }
                num_k := nums[k]
                for k < len(nums) && num_k == nums[k] {
                    k++
                }
                l = k + 1
            }
            num_j := nums[j]
            for j < len(nums) && num_j == nums[j] {
                j++
            }
	        k = j + 1
	        l = k + 1
        }
        num_i := nums[i]
        for i < len(nums) && num_i == nums[i] {
            i++
        }
	    j = i + 1
	    k = j + 1
	    l = k + 1
    }
    
    return result;
}
