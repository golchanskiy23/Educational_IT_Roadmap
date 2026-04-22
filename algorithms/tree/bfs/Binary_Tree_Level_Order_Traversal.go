/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var v [][]int
    f(root, 0, &v)
    return v
}

func f(root *TreeNode, depth int, v *[][]int){
    if root == nil{return}
    if root.Left == nil && root.Right == nil{
        if depth >= len(*v){
            tmp := []int{root.Val}
            *v = append(*v, tmp)
        } else{
            (*v)[depth] = append((*v)[depth], root.Val)
        }
        return
    }

    if depth >= len(*v){
        tmp := []int{root.Val}
        *v = append(*v, tmp)
    } else{
        (*v)[depth] = append((*v)[depth], root.Val)
    }

    f(root.Left, depth+1, v)
    f(root.Right, depth+1, v)   
}