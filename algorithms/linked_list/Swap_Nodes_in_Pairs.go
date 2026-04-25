/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    prev := dummy

    for prev.Next != nil && prev.Next.Next != nil{
        f := prev.Next
        s := prev.Next.Next

        f.Next = s.Next
        s.Next = f
        prev.Next = s
        prev = f
    }

    return dummy.Next
}