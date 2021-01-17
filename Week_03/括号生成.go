/**
  @author: dingfeng
  @date: 2021/1/15
  @note:
**/
package Week_03

func GenerateParenthesis(n int) []string {
	res := []string{}
	generatePar(0, 0, n, "", &res)
	return res
}

// APPEND的一个大坑 切片不一定是原引用 可以用*来解决
func generatePar(left, right, n int, s string, ss *[]string) {
	//terminator
	if left == n && right == n {
		*ss = append(*ss, s)
		return
	}
	if left < n {
		//drill down
		generatePar(left+1, right, n, s+"(", ss)
	}
	if right < left {
		//drill down
		generatePar(left, right+1, n, s+")", ss)
	}
	//reverse state
}

var res []string
func GenerateParenthesis2(n int) []string {
	generatePar2(0, 0, n, "")
	return res
}
func generatePar2(left, right, n int, s string) {
	if left == n && right == n {
		res = append(res, s)
		return
	}
	if left < n {
		generatePar2(left+1, right, n, s+"(")
	}
	if right < left {
		generatePar2(left, right+1, n, s+")")
	}
}

func GenerateParenthesis3(n int) []string {
	res := []string{}
	var dfs func(left, right, n int, s string)
	dfs = func(left, right, n int, s string){
		if left == n && right == n {
			res = append(res, s)
			return
		}
		if left < n {
			dfs(left+1, right, n, s+"(")
		}
		if right < left {
			dfs(left, right+1, n, s+")")
		}
	}
	dfs(0, 0, n, "")
	return res
}
