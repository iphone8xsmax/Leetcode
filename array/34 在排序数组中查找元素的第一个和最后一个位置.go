给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 改进的二分搜索
func searchRange(nums []int, target int) []int {
	l := edgeBinSearch(nums, true, target)
	r := edgeBinSearch(nums, false, target)
	return []int{l, r}
}

// 普通二分搜索在匹配到 target 时直接 return
// 在本题搜索时在匹配到 target 之后依旧向边缘走当做没匹配到，注意 2 个边界条件
// O(logN) // ok
func edgeBinSearch(nums []int, leftest bool, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case target < nums[mid]:
			r = mid - 1
		case target > nums[mid]:
			l = mid + 1
		default:
			if leftest {
				if mid == 0 || nums[mid] > nums[mid-1] { // 不再继续向左走的 2 个边界条件
					return mid
				}
				r = mid - 1 // 继续在左侧找
			} else {
				if mid == n-1 || nums[mid] < nums[mid+1] {
					return mid
				}
				l = mid + 1 // 继续在右侧找
			}
		}
	}
	return -1
}