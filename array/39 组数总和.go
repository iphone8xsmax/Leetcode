给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func combinationSum(candidates []int, target int) [][]int {
	size := len(candidates)  //长度
	var res [][]int
	var path []int
	sort.Ints(candidates)  //排序
	Dfs(path, candidates, 0,size, &res, target)
	return res
}

func Dfs(path, candidates []int, begin int, size int, res *[][]int, target int){
	if target == 0{  //得到结果
		tmp := make([]int, len(path))  //临时切片
		copy(tmp, path)  //保存一组符合要求的数据
		*res = append(*res, tmp)  //返回结果
		return
	}

	//将每一个可能符合条件的数据放入，然后得到新的target，再进行递归，直到发现无法满足为止
	//因为可以重复，所以for循环进行多次遍历，否则需要判别是否重复，continue
	for index:=begin;index<size;index++{
		residue := target - candidates[index] 
		if residue < 0{
			break
		}
		path = append(path, candidates[index])

		Dfs(path, candidates, index, size, res, residue)
		path = path[0:len(path)-1]
	}
}