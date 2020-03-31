func twoSum(nums []int, target int) []int {
    num2 := make(map[int]int,len(nums))
    for i, num := range nums{
        pair := target - num
        if j, ok := num2[pair]; ok && i != j{
            return []int{j ,i}
        }
        num2[num] = i
    }
    return nil
}