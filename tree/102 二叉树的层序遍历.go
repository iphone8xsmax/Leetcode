给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil{
        return nil
    }

    var res [][]int
    queue := []*TreeNode{root} //创建队列，用于存储节点信息，首先存入根节点，之后存入左和右节点
    for len(queue) > 0{
        var floor []int
        level := len(queue) //当前层的节点数
        for i := 0; i < level; i++{
            cur := queue[0]  //取出第一个节点
            queue = queue[1:] //队列更新为取出之后的状态
            if cur.Left != nil{  //先加入左节点
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil{
                queue = append(queue, cur.Right)
            }
            floor = append(floor, cur.Val) //存储值
        }
        res = append(res, floor) //将对应层的节点数据依次加入
    }
    return res
}