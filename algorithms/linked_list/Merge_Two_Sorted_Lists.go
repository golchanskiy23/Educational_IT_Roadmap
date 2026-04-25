/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    curr := &ListNode{0, nil}
    start := curr
    for list1 != nil && list2 != nil{
        if list1.Val <= list2.Val{
            start.Next = list1
            list1 = list1.Next
            start = start.Next
        } else{
            start.Next = list2
            list2 = list2.Next
            start = start.Next
        }
    }

    for list1 != nil{
        start.Next = list1
        list1 = list1.Next
        start = start.Next
    }

    for list2 != nil{
        start.Next = list2
        list2 = list2.Next
        start = start.Next
    }

    return curr.Next
}