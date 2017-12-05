/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    less_vals := make([]int, 0)
    equal_vals := make([]int, 0)
    
    current := head
    for ; current != nil; current = current.Next {
        if current.Val < x {
            less_vals = append(less_vals, current.Val)
        } else {
            equal_vals = append(equal_vals, current.Val)
        }
    }
    
    result := new(ListNode)
    current = result
    for i := 0; i < len(less_vals); i++ {
        current.Next = new(ListNode)
        current.Next.Val = less_vals[i]
        current = current.Next
    }
    for i := 0; i < len(equal_vals); i++ {
        current.Next = new(ListNode)
        current.Next.Val = equal_vals[i]
        current = current.Next
    }
    
    return result.Next
}
