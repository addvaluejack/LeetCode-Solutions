/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    pre_head := &ListNode{0, head}
    wait_for_reverse := make([]*ListNode, 0)
    
    if k == 0 || k == 1 || head == nil{
        return head
    }
    
    foo(pre_head, k, wait_for_reverse, 0)
    
    return pre_head.Next
}

func foo(current *ListNode, k int, wait_for_reverse []*ListNode, acc int) {
    if acc == k {
        if wait_for_reverse[acc-1].Next != nil {
            another_wait_for_reverse := make([]*ListNode, 0)
            foo(wait_for_reverse[acc-1], k, another_wait_for_reverse, 0)
        }
        wait_for_reverse[0].Next = wait_for_reverse[acc-1].Next
        current.Next = wait_for_reverse[acc-1]
        for i := acc - 1; i > 0; i-- {
            wait_for_reverse[i].Next = wait_for_reverse[i-1]
        }
    }
    if acc < k {
        if acc == 0 && current.Next != nil{
            wait_for_reverse = append(wait_for_reverse, current.Next)
            foo(current, k, wait_for_reverse, 1)
        } else if wait_for_reverse[acc-1].Next != nil {
            wait_for_reverse = append(wait_for_reverse, wait_for_reverse[acc-1].Next)
            foo(current, k, wait_for_reverse, acc+1)
        }
    }
}
