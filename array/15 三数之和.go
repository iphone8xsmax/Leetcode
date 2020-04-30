给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

//输出结果和原始顺序无关，则可以先排序，方便去重，头尾结合的双指针法
import "sort"
func threeSum(nums []int) [][]int {
    sort.Ints(nums) //排序
    n := len(nums)
    var res [][]int
    //计算的是从i向后的
    for i, num := range nums{
        if num > 0{ //nums[i]大于0了，之后的肯定也大于0
            break
        }
        if i > 0 && nums[i] == nums[i-1]{ //向前去一次重，否则可能会继承前一次的结果，出现重复=0的数组
            continue
        }

        l := i + 1
        r := n - 1
        for l < r{
            sum := num + nums[l] + nums[r]
            switch{
            case sum > 0:
                r--
            case sum < 0:
                l++
            default:
                res = append(res, []int{num, nums[l], nums[r]})

                //第二层候选数向后去一次重，防止移动l,r指针后两个数组一样
                for l < r && nums[l] == nums[l+1]{
                    l++
                }
                for r > l && nums[r] == nums[r-1]{
                    r--
                }
                l++
                r--
            }
        }

    }

    return res
}
