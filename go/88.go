func merge(nums1 []int, m int, nums2 []int, n int)  {
    j := m-1
    k := n-1
    
    for i := m+n-1; i >= 0; i-- {
        if k < 0 {
            break
        }
        if j < 0 {
            for ;i >=0; i-- {
                nums1[i] = nums2[k]
                k--
            }
            break
        }
        if nums1[j] > nums2[k] {
            nums1[i] = nums1[j]
            j--
        } else {
            nums1[i] = nums2[k]
            k--
        }
    }
}
