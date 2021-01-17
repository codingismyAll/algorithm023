/**
  @author: dingfeng
  @date: 2021/1/15
  @note:
**/
package Week_03

//动态规划
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
// 动态规划优化(节约内存)
func climbStairs2(n int) int {
	p:=0;q:=0;r:=1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p+q
	}
	return r
}
// 递归
func climbStairs3(n int) int {
	if n < 3{
		return n
	}
	return climbStairs3(n-1)+climbStairs3(n-2)
}
// 递归转尾递归
func climbStairs4(n int) int {
	return Fibonacci(n, 1, 1)
}
func Fibonacci (n,a,b int)int  {
	if n <= 1 {
		return b
	}
	return Fibonacci(n - 1, b, a + b)
}