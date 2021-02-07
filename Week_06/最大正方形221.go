/**
  @author: dingfeng
  @date: 2021/2/6
  @note:
**/
package leetcode

/*
暴力求解
时间复杂度：O(mnmin(m,n)^2)，其中 m 和 n 是矩阵的行数和列数。

需要遍历整个矩阵寻找每个 1，遍历矩阵的时间复杂度是 O(mn)。
对于每个可能的正方形，其边长不超过 m 和 n 中的最小值，需要遍历该正
方形中的每个元素判断是不是只包含 1，遍历正方形时间复杂度是 O(mn min(m,n)^2)

总时间复杂度是 O(mn min(m,n)^2)

空间复杂度：O(1)。额外使用的空间复杂度为常数。

*/
func maximalSquare(matrix [][]byte) int {

	row := len(matrix)
	if row == 0 {
		return 0
	}
	column := len(matrix[0])
	if column == 0 {
		return 0
	}
	maxside := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == '1' {
				maxside = max(maxside, 1)
				currmaxside := min(row-i, column-j)
				for k := 1; k < currmaxside; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+m][j+k] == '0' || matrix[i+k][j+m] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxside = max(maxside, k+1)
					} else {
						break
					}

				}
			}
		}

	}

	return maxside * maxside
}

/*
可以用动态规划来解决
if a[i][j] == '1'
dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
if a[i][j] == '0'
dp[i][j] = 0
dp[i][j]表示以i和j为右下角组成的最大正方形
*/
func maximalSquare2(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	column := len(matrix[0])
	if column == 0 {
		return 0
	}

	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, column)
	}
	maxside := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				if i == 0 {
					dp[i][j] = 1
				} else if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				}
			}
			maxside = max(maxside, dp[i][j])
		}
	}
	return maxside * maxside

}
