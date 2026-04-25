/*
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head.Next == nil{return head}
    turtle, hare := head, head.Next
    for hare != nil{
        if hare.Next == nil{
            return turtle.Next
        }
        turtle = turtle.Next
        hare = hare.Next.Next
    }
    return turtle 
}