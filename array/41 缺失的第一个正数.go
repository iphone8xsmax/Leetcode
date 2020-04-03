给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 对数组进行归位预处理
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			swap(&nums[i], &nums[nums[i]-1])
		}
	}

	// 向后检查
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 理想正整数数组
	return n + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}