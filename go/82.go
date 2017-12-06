/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    fake_head := new(ListNode)
    fake_head.Next = head
    current := fake_head
    
    for ; current.Next != nil;  {
        if current.Next.Next != nil {
            tmp := current.Next.Next
            if tmp.Val == current.Next.Val {
                for ; tmp != nil && tmp.Val == current.Next.Val; tmp = tmp.Next {
                }
                current.Next = tmp
            } else {
                current = current.Next
            }
        } else {
            break
        }
    }
    
    return fake_head.Next
}
