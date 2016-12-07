/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var result, current, tmp *ListNode
    c1, c2 := l1, l2
    
    for c1 != nil && c2 != nil {
        if result == nil {
            if c1.Val < c2.Val {
                result = c1
                current = result
                c1 = c1.Next
            } else {
                result = c2
                current = result
                c2 = c2.Next
            }
        } else {
            if c1.Val < c2.Val {
                tmp = c1
                c1 = c1.Next
                current.Next = tmp
                current = current.Next
            } else {
                tmp = c2
                c2 = c2.Next
                current.Next =tmp
                current = current.Next
            }
        }
    }
    
    if c1 != nil {
        if result == nil {
            return c1
        }
        current.Next = c1
    }
    if c2 != nil {
        if result == nil {
            return c2
        }
        current.Next = c2
    }
    
    return result
}
