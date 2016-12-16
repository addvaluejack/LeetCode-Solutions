/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    fake_head := ListNode{0, head}
    current := &fake_head
    
    for {
        if current.Next == nil {
            break
        } else {
            if current.Next.Next == nil {
                break
            } else {
                tmp0 := current.Next
                tmp1 := current.Next.Next.Next
                current.Next = current.Next.Next
                current.Next.Next = tmp0
                current.Next.Next.Next = tmp1
                current = current.Next.Next
            }
        }
    }
    
    return fake_head.Next
}
