package function

import "sort"

func test(mat [][]int, k int) []int {

	result := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				result[i] += mat[i][j] // 获取这一行的值
			} else {
				break
			}
		}
	}
	// 找到result中最小的k个数：最好使用的是最小堆，但是Go没有:使用三个数组，首先是排序，找到应该最小的几个值，然后是重新去原数组找到这个值所在的位置，最后是记录这个位置索引，使用第三个数组表示是否被使用过
	// sorts排序时，相等值的位置不变
	result2 := make([]int, len(result))
	copy(result2, result)
	sort.Ints(result2)
	result3 := make([]int, len(result)) // 表示是否取出
	result4 := make([]int, k)
	for i := 0; i < k; i++ { //取k个最小值
		for j := 0; j < len(result); j++ {
			if result2[i] == result[j] && result3[j] == 0 {
				result3[j] = 1
				result4[i] = j //记录这个结果
				break
			}
		}
	}
	return result4
}
