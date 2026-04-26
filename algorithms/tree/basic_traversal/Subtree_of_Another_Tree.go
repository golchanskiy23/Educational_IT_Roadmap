/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(root, sub *TreeNode) bool{
    if root == nil && sub == nil {return true}
    if root == nil || sub == nil {return false}
    return root.Val == sub.Val && isSameTree(root.Left, sub.Left) && isSameTree(root.Right, sub.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {return false}
    if isSameTree(root, subRoot) {
        fmt.Println(root.Val)
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}