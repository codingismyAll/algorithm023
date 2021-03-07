/**
  @author: dingfeng
  @date: 2021/1/25
  @note:
**/
package Week_05

func foo(key, val int, store map[int]int) int{
	if v, ok := store[key]; ok {
		if val == v {
			return 1
		}
		return 2
	}else{
		store[key] = val
		return 0
	}
}
