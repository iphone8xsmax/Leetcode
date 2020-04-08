给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种


//G(n)= 
i=1
∑
n
​	
 G(i−1)⋅G(n−i)  卡塔兰数   Cn+1 = 2(2n+1)/(n+2) * Cn


func numTrees(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = res * 2 * (2*i + 1) / (i + 2)
	}
	return res
}