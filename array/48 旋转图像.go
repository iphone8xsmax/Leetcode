给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rotate(matrix [][]int) [][]int { //正方形的旋转可以改为多重的折叠
    n := len(matrix)
    if n <= 1{
        return matrix
    }
    //右侧对角线折叠，左上右下
    for i := 0; i < n; i++{
        for j := i + 1; j < n; j++{
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    //左右折叠
    for i := 0; i < n; i++{
        for j := 0; j < n/2; j++{
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }

    return matrix
}
