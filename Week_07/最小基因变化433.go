/**
  @author: dingfeng
  @date: 2021/1/27
  @note:
**/
package leetcode

func minMutation(start string, end string, bank []string) int {
	bankmap := make(map[string]bool)
	for _, item := range bank {
		bankmap[item] = true
	}
	if _, ok := bankmap[end]; !ok {
		return -1
	}
	four := []byte{'A', 'C', 'G', 'T'}
	step := 0
	queue := []string{start}
	for len(queue) > 0 {
		step++
		n := len(queue)
		for i := 0; i < n; i++ {
			currstr := queue[i]
			for j := 0; j < len(currstr); j++ {
				//oldchar := currstr[j]
				for k := 0; k < len(four); k++ {
					var transtr string
					if j+1 == len(currstr) {
						transtr = currstr[:j] + string(four[k])
					} else {
						transtr = currstr[:j] + string(four[k]) + currstr[j+1:]
					}
					if transtr == end {
						return step
					}
					if _, ok := bankmap[transtr]; ok {
						queue = append(queue, transtr)
						delete(bankmap, transtr)
					}
				}
			}
		}
		//出对列
		queue = queue[n:]
	}
	return -1
}
