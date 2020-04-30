
//双指针法
// 两条线段之间的面积受限与最短的线段，线段间距越长，面积越大
// 使用 2 个指针指向首部和尾部，将短指针向长指针方向移动，看能不能找到更长的线，使面积更大
// 依据：向长线方向每次移动 1 格，虽然宽度-1，但是(高度变高)*(宽度-1) >= 高度*宽度
func maxArea(height []int) int {
    maxS := 0
    left := 0
    right := len(height) - 1

    for left < right {
        width := right - left
        heigth := min(height[left], height[right])
        maxS = max(maxS, width * heigth)

        if height[left] <= height[right]{
            left++
        }else {
            right--
        }
    }
    return maxS
}

func min(num ...int)int{
    if len(num) == 0{
        return -1
    }
    m := num[0]
    for _, n := range num{
        if n < m{
            m = n
        }
    }
    return m
}


func max(num ...int)int{
    if len(num) == 0{
        return -1
    }
    m := num[0]
    for _, n := range num{
        if n > m{
            m = n
        }
    }
    return m
}
