/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func size(head *ListNode) int{
    curr := head
    sz := 0
    for curr != nil{
        curr = curr.Next
        sz++
    }
    return sz
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    k := size(head)
    if k == 1{return nil}
    if n == k{return head.Next}

    idx := 1
    curr := head
    for idx < size(head)-n{
        curr = curr.Next
        idx++
    }

    next := curr.Next.Next
    curr.Next = next

    return head
}