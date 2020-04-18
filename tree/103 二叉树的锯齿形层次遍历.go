给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res = [][]int{}
    
	var queue = []*TreeNode{root}
	var level, counter, knife, cur int
	for len(queue) > 0 {
		counter = len(queue)
		knife = counter
		res = append(res, []int{})
		for 0 < counter {
			cur = knife - counter
			if queue[cur].Left != nil {
				queue = append(queue, queue[cur].Left)
			}
			if queue[cur].Right != nil {
				queue = append(queue, queue[cur].Right)
			}

			counter--
			if level%2 != 0 {
				cur = counter
			}
			res[level] = append(res[level], queue[cur].Val)
		}
		queue = queue[knife:]
		level++
	}
	return res
}