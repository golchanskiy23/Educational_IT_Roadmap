/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.
Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
*/

/*
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    idx int    
}

func Constructor() Codec {
   return Codec{
    idx: 0,
   } 
}

func preorder_serialize(root *TreeNode, v *[]string){
    if root == nil{
        *v = append(*v, "null")
        return
    }
    
    *v = append(*v, strconv.Itoa(root.Val))
    preorder_serialize(root.Left, v)
    preorder_serialize(root.Right, v)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var v []string
    preorder_serialize(root, &v)
    return strings.Join(v , ",")
}

func (this *Codec) preorder_deserialize(arr []string) *TreeNode{
    if arr[this.idx] == "null"{
        this.idx++
        return nil
    }

    num , _ := strconv.Atoi(arr[this.idx])
    this.idx++

    root := &TreeNode{Val: num,}

    root.Left = this.preorder_deserialize(arr)
    root.Right = this.preorder_deserialize(arr)

    return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    arr := strings.Split(data, ",")
    this.idx = 0
    return this.preorder_deserialize(arr)
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
*/