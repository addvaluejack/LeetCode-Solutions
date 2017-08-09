/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    
    if head == nil {
        return head
    }
    
    for ;; {
        if current.Next == nil {
            break
        }
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
            continue
        }
        current = current.Next
    }
    
    return head
}
