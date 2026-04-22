/*
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// iterative_BFS
func connect(root *Node) *Node {
    if root == nil{return nil}
    queue := []*Node{root}
    for len(queue) > 0{
        var curr *Node
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            if front.Left != nil{queue = append(queue, front.Left)}
            if front.Right != nil{queue = append(queue, front.Right)}

            if curr != nil{
                curr.Next = front
            }
            curr = front
        }
        curr.Next = nil
    }
    return root
}