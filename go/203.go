/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    
    current := head
    
    for ; current != nil && current.Val == val; {
        head = head.Next
        current = current.Next
    }
    
    for ; current != nil; {
        if current.Next == nil {
            break
        }
        if current.Next.Val == val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    
    return head
}
