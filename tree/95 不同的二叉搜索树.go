给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0{
        return nil
    }
    return generate(1, n)
}

func generate(start, end int) []*TreeNode{
    var roots []*TreeNode

    switch{
    case start > end: //某一边子树无节点，左斜树或右斜树
        roots = append(roots, nil)
    case start == end: //某一子树只有一个节点
        roots = append(roots, &TreeNode{Val: start})
    case start < end:
        for rootV := start; rootV <= end; rootV++{
            leftTrees := generate(start, rootV-1)
            rightTrees := generate(rootV+1, end)

        //组合左右子树
        for _, ltree := range leftTrees{
            for _, rtree := range rightTrees{
                root := &TreeNode{
                    Val: rootV,
                    Left: ltree,
                    Right: rtree,
                }
                roots = append(roots, root)
            }
        }
        }
    }
    return roots
}