给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。


递归：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return mirror(root, root)
}

func mirror(node1, node2 *TreeNode) bool{
    if node1 == nil && node2 == nil{
        return true
    }
    if node1 == nil || node2 == nil{
        return false
    }
    return node1.Val == node2.Val && mirror(node1.Left, node2.Right) && mirror(node1.Right, node2.Left)
}



迭代：使用队列实现
func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) >= 2 {
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		queue = append(queue, t1.Left)
		queue = append(queue, t2.Right)
		queue = append(queue, t1.Right)
		queue = append(queue, t2.Left)
	}
	return true
}
