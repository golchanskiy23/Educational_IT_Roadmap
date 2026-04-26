/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // оба поддерева - не nil - значит root - LCA
    if root == nil {return nil}

    if root == p || root == q{
        return root
    }

    l := lowestCommonAncestor(root.Left, p,q)
    r := lowestCommonAncestor(root.Right, p,q)

    if l == nil && r != nil{return r}
    if r == nil && l != nil{return l}
    if r == nil && r == nil{return nil}

    return root
}