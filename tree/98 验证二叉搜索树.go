给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil{
        return true
    }
    return verity(root, -1<<63, 1<<63-1)
}

func verity(node *TreeNode, min, max int) bool{
    if node == nil{  
        return true
    }
    if node.Val >= max || node.Val <= min{  //用于判定是否满足，每次更新最大最小值，左树右树要分别满足
        return false
    }
    //左树的左子节点要满足大于最小值，小于父节点值，右树右子节点值要大于父节点，小于最大值
    return verity(node.Left, min, node.Val) && verity(node.Right, node.Val, max)
}