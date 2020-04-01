给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。


func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) //排序
	var res [][]int 
	nSum(nums, target, 0, 4, &res, []int{}) //调用求和数组函数
	return res
}

//传入数组，目标值，左右指针值，结果数组集，
//相当于，先把有可能满足的两个数存入数组，再用双指针法找其余两个数
func nSum(nums []int, target int, pos int, n int, res *[][]int, cur []int) {
	if n == 2 {
		left, right := pos, len(nums) - 1
		for left < right {
			if sum := nums[left] + nums[right]; sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				t := make([]int, len(cur) + 2) //如果满足条件了。创建一个比先找的几位数多两位即指针的切片，把这两个数放进去
				copy(t, cur)
				t[len(t) - 2] = nums[left]
				t[len(t) - 1] = nums[right]
				*res = append(*res, t)
				left++
				right--
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
				for left < right && nums[right] == nums[right + 1] {
					right--
				}
			}
		}
		return
	}

	for i := pos; i < len(nums) - n + 1; i++ {
		if target < nums[i] * n || target > nums[len(nums)-1] * n { //优化，之后的循环不可能满足了
			break
		}
		if i > pos && nums[i] == nums[i-1] { //去重复
			continue
		}
		cur = append(cur, nums[i])  //把可能满足的元素加入cur
		nSum(nums, target - nums[i], i + 1, n - 1, res, cur)  //递归调用
		cur = cur[:len(cur)-1] //移动一位
	}
}