给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
   return same(p, q)
}

func same(node1, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil{
        return true
    }
    if node1 == nil || node2 == nil{
        return false
    }

    return node1.Val == node2.Val && same(node1.Left, node2.Left) && same(node1.Right, node2.Right)
}