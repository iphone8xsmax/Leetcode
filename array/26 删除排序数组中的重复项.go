//给定一个排序数组，你需要在 原地 //删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) //额外空间的条件下完成。

func removeDuplicates(nums []int) int {
n := len(nums)
if n < 2{ //特殊
return n
}
l, r := 0, 1
for r < n{
if nums[l] < nums[r]{ //利用数组有序，未判别的大于等于之前的
l++
nums[l] = nums[r]
}
r++
}
return l + 1
}

作者：mojinuo
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/solution/8ms-shi-jian-9430-jing-dian-shuang-zhi-zhen-fa-by-/
来源：力扣（LeetCode）
