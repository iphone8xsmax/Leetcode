// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n1 + n2 == 0{
        return -1
    }
    if (n1 + n2) % 2 == 0{
        l := findK(nums1, nums2, (n1 + n2) / 2)
        r := findK(nums1, nums2, (n1 + n2) / 2 +1)
        return float64(l + r) / 2
    }

    return float64(findK(nums1, nums2, (n1 + n2) / 2 + 1))
}
//寻找第K的数
func findK(nums1 []int, nums2 []int, k int) int{
    n1 := len(nums1)
    n2 := len(nums2)
    if n1 > n2{ //将短数组放在前面
        n1, n2 = n2, n1
        nums1, nums2 = nums2, nums1
    }
    //处理特殊情况
    if n1 == 0{
        return nums2[k-1]
    }
    if k == 1{
        return min(nums1[0], nums2[0])
    }

    k1 := min(k/2, n1)//避免越界
    k2 := k - k1

    switch{ //递归
        case nums1[k1-1] < nums2[k2-1]:
            return findK(nums1[k1:], nums2, k2)
        case nums1[k1-1] > nums2[k2-1]:
            return findK(nums1, nums2[k2:], k1)
        default:
            return nums1[k1-1]
    }
}

func min(a int, b int) int{
    switch{
        case a > b:
        return b
        case a < b:
        return a
        default:
        return a
    }
}