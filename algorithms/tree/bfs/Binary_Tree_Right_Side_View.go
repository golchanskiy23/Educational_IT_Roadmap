/*
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var v []int
    f(root, 0, &v)
    return v
}

func f(root *TreeNode,depth int, v *[]int){
    if root == nil{return}
    if root.Left == nil && root.Right == nil{
        if depth >= len(*v){
            *v = append(*v, root.Val)
        } else{
            (*v)[depth] = root.Val
        }
        return
    }

    if depth >= len(*v){
        *v = append(*v, root.Val)
    } else{
        (*v)[depth] = root.Val
    }
    f(root.Left, depth+1, v)
    f(root.Right, depth+1, v)
}
*/