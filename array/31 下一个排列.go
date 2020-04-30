实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation

func nextPermutation(nums []int) {
	length := len(nums)
	exist := false
	i := length - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			exist = true
			break
		}
	}

	if !exist { //nums降序排列，改为升序
		sort.Ints(nums)
	} else { //找到了需要调整数序的位置，则在i后面找出大于nums[i]的最小数和i-1进行交换，再将i后面的数进行升序排序
		min := nums[i]
		k := i
		for j := i; j < length; j++ {
			if nums[j] < min && nums[j] > nums[i-1] {
				min = nums[j]
				k = j
			}
		}
		nums[k] = nums[i-1]
		nums[i-1] = min
		sort.Ints(nums[i:])

	}
}
