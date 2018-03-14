/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func foo(aPtr **ListNode, b *ListNode, lengthPtr *int, countPtr *int) {
    (*lengthPtr)++
    if b.Next.Next == nil {
        (*lengthPtr)++
        tmp := (*aPtr).Next
        (*aPtr).Next = b.Next
        b.Next.Next = tmp
        *aPtr = tmp
        (*countPtr)++
    } else {
        foo(aPtr, b.Next, lengthPtr, countPtr)
        if (*countPtr) == (*lengthPtr-1)/2 {
            if *lengthPtr%2 == 0 {
                (*aPtr).Next.Next = nil
            } else {
                (*aPtr).Next = nil
            }
        } else {
            tmp := (*aPtr).Next
            (*aPtr).Next = b.Next
            b.Next.Next = tmp
            *aPtr = tmp
            (*countPtr)++
        }

    }
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return
    }
    length := 0
    count := 0
    foo(&head, head, &length, &count)
}
