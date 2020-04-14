二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//使用中序遍历的方法，遇到错误节点则交换
var last, first, second *TreeNode
func recoverTree(root *TreeNode)  {
    last, first, second = nil, nil, nil
    dfs(root)
    first.Val, second.Val = second.Val, first.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left) //找到最下层 
	if last != nil && root.Val <= last.Val {
		if first != nil {
			second = root
			return //剪枝
		}
		first, second = last, root
	}
	last = root //last
	dfs(root.Right)
}
