/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// what if often modified ?
func inorder(root *TreeNode, k,res *int){
    if root == nil{return}

    inorder(root.Left, k,res)
    (*k)--
    if (*k) == 0 {
        *res = root.Val
        return
    }
    inorder(root.Right, k,res)
}

func return_inorder(root *TreeNode, k *int) (bool, int){
    if root == nil{return false,0}

    if ok, v := return_inorder(root.Left, k); ok{
        return true, v
    }

    (*k)--
    if (*k) == 0 {
        return true, root.Val
    }

    return return_inorder(root.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
    /*var res int
    inorder(root, &k, &res)*/
    _, res := return_inorder(root, &k)
    return res
}