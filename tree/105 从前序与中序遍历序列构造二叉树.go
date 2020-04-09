根据一棵树的前序遍历与中序遍历构造二叉树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode{
    if len(preorder) == 0 || len(inorder) == 0{
        return nil
    }
    if len(preorder) == 1{
        return &TreeNode{Val:preorder[0]}
    }
    if len(inorder) == 1{
        return &TreeNode{Val:inorder[0]}
    }
    root := &TreeNode{Val:preorder[0]}  //构建的根节点为前序遍历的第一个元素

    var i int
    for ; i < len(inorder); i++{
        if inorder[i] == preorder[0]{  //中序遍历到根节点了，代表左子树已经遍历完
            break
        }
    }
    lPreorder := preorder[1:i+1]   //左子树是前序遍历到n+1元素为止
    rPreorder := preorder[i+1:]

    //递归，分别构建左子树以及右子树
    root.Left = build(lPreorder, inorder[:i]) 
    root.Right = build(rPreorder, inorder[i+1:])
    return root
}