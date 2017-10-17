/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }
    
    length := 0
    
    for current := head; current != nil; current = current.Next {
        length++
    }
    
    if k >= length {
        k = k%length
    }
    
    if k == 0 {
        return head
    }
    
    x := head
    y := head
    var result *ListNode
    
    for i := 0; i < k; i++ {
        y = y.Next
    }
    
    for ; y.Next != nil; {
        x = x.Next
        y = y.Next
    }
    
    result = x.Next
    x.Next = nil
    y.Next = head
    
    return result
}
