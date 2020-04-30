给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

//经典二分法
func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    if right<0 || nums[0]>target{
        return 0
    }
    if nums[right]<target {
        return right+1
    }
    for left <=right{
        mid := (left+right)/2
        if nums[mid]==target{
            return mid
        }
        if nums[mid]<target{
            left = mid+1
        }
        if nums[mid]>target{
            right = mid-1
        }
    }
    return left
}
