94.给定一个二叉树，返回它的中序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int
    traverse(root, &res)
    return res
}

func traverse(node *TreeNode, res *[]int){
    if node == nil{
        return
    }
    if node.Left != nil{
        traverse(node.Left, res)
    }
    *res = append(*res, node.Val)
    if node.Right != nil{
        traverse(node.Right, res)
    }
}