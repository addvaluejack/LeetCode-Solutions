import "math/rand"

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func foo(result_ptr *[][]int, nums []int, mark []int, index int, tmp []int) {
    if index == len(nums)-1 {
        permutation := make([]int, len(nums))
        for i := 0; i < len(nums); i++ {
            permutation[i] = nums[mark[i]];
        }
        (*result_ptr) = append((*result_ptr), permutation);
        return
    }
    for i := 0; i < len(nums); {
        if tmp[i] == 1 {
            i++
            continue
        }
        new_tmp := make([]int, len(nums))
        for k := 0; k < len(nums); k++ {
            new_tmp[k] = tmp[k]
        }
        new_mark := make([]int, len(nums))
        for k := 0; k < len(nums); k++ {
            new_mark[k] = mark[k]
        }
        new_mark[index+1] = i
        new_tmp[i] = 1
        foo(result_ptr, nums, new_mark, index+1, new_tmp)
        j := i+1
        for ; j < len(nums) && nums[j] == nums[i]; j++ {
            
        }
        i = j
    }
}

func permuteUnique(nums []int) [][]int {
    nums = QuickSort(nums)
    var result [][]int
    
    if len(nums) == 0 {
        var tmp []int
        return append(result, tmp)
    }
    
    mark := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        mark[i] = -1
    }
    
    for i := 0; i < len(nums); {
        tmp := make([]int, len(nums))
        tmp[i] = 1
        mark[0] = i
        foo(&result, nums, mark, 0, tmp)
        j := i+1
        for ; j < len(nums) && nums[j] == nums[i]; j++ {
            
        }
        i = j
    }
    
    return result
}
