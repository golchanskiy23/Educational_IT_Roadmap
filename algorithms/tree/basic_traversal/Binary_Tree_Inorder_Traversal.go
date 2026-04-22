/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    var v []int
    f(root, &v)
    return v
}

func f(root *TreeNode, v *[]int){
    if root == nil{return}
    if root.Left == nil && root.Right == nil{
        *v = append(*v, root.Val)
        return 
    }

    f(root.Left, v)
    *v = append(*v, root.Val)
    f(root.Right, v)
}