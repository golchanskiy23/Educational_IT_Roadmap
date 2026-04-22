/*
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var prev *TreeNode

func flatten(root *TreeNode)  {
    prev = nil
    rebuild(root)
}

func rebuild(root *TreeNode){
    if root == nil {return}
    
    rebuild(root.Right)
    rebuild(root.Left)

    root.Right = prev
    root.Left = nil
    prev = root
}