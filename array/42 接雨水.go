
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。


示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

func trap(height []int) int {
    area := 0
    l, r := 0, len(height) - 1
    lTop, rTop := 0, 0

    for l < r{
        lTop = max(height[l], lTop) //动态存储左右最高点
        rTop = max(height[r], rTop)

        if lTop < rTop{ //在当前高度小于最高高度时，可以积累雨水。哪边高往那边移，必有柱子可以挡水，不需要计算考虑
            area += lTop - height[l] //右侧更高点，往右侧移，计算高度差累计面积
            l++
        }else{
            area += rTop - height[r] //向左侧移，用右侧最高高度来计算高度差，累计面积
            r--
        }
    }
    return area
}

func max(x, y int) int {
    if x > y{
        return x
    }
    return y
}