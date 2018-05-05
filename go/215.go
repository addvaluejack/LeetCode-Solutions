func foo(k_array []int, start int, end int, target int) int {
    if k_array[(start+end)/2] == target {
        return (start+end)/2
    }
    if start == end {
        if k_array[(start+end)/2] > target {
            return (start+end)/2+1
        } else {
            return (start+end)/2
        }
    } else {
        if target > k_array[(start+end)/2] {
            return foo(k_array, start, (start+end)/2, target)
        } else {
            return foo(k_array, (start+end)/2+1, end, target)
        }
    }
}

func findKthLargest(nums []int, k int) int {
    if k <= 0 || k > len(nums) {
        return -1
    }
    k_array := make([]int, k)
    k_array_size := 0
    
    for i := 0; i < len(nums); i++ {
        if k_array_size < k {
            j := 0
            for ; j < k_array_size; j++ {
                if nums[i] >= k_array[j] {
                    break
                }
            }
            k_array = append(k_array[0:j], append([]int{nums[i]}, k_array[j:len(k_array)-1]...)...)
            k_array_size++
        } else {
            if nums[i] > k_array[k-1] {
                index := foo(k_array, 0, k-2, nums[i])
                k_array = append(k_array[0:index], append([]int{nums[i]}, k_array[index:len(k_array)-1]...)...)
            }
        }
        
    }
    
    return k_array[k-1]
}
