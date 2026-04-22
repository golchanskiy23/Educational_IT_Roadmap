/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

-The left subtree of a node contains only nodes with keys strictly less than the node's key.
-The right subtree of a node contains only nodes with keys strictly greater than the node's key.
-Both the left and right subtrees must also be binary search trees.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// CHECK GLOBALLY
func validate(root *TreeNode, min, max *int) bool{
    if root == nil {return true}

    if min != nil && root.Val <= *min{
        return false
    }   

    if max != nil && root.Val >= *max{
        return false
    }

    return validate(root.Left, min, &root.Val) && 
        validate(root.Right, &root.Val, max)
}
 
func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}