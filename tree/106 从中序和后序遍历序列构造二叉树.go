根据一棵树的中序遍历与后序遍历构造二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    return build(inorder, postorder)
}

func build(inorder []int, postorder []int) *TreeNode{
    if len(inorder) == 0 || len(postorder) == 0{
        return nil
    }
    if len(inorder) == 1{
        return &TreeNode{Val:inorder[0]}
    }
    if len(postorder) == 1{
        return &TreeNode{Val:postorder[0]}
    }

    //后序遍历最后一个节点是根节点
    rootV := postorder[len(postorder)-1]
    var i int 
    for ; i < len(inorder); i++{
        if inorder[i] == rootV{
            break
        }
    }
    root := &TreeNode{Val:rootV}
    root.Left = build(inorder[:i], postorder[:i]) //left tree
    root.Right = build(inorder[i+1:], postorder[i:len(postorder)-1]) //right tree

    return root
}