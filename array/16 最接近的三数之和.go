给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

// 双指针法
// 和 15 一样，排序预处理能知道双指针移动的方向，记录最小 abs
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := 1<<31 - 1
	minSum := 0

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 优化：可选的去重
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if abs(target-sum) < minAbs {
				minAbs = abs(target - sum)
				minSum = sum
			}
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				return target
			}
		}
	}
	return minSum
}
