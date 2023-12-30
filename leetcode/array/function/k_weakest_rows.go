package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {

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

func Exec1337() {
	// 填充式输出题目名称
	title := "矩阵中战斗力最弱的 K 行"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var mat [][]int
	var k int
	var str1 string
	fmt.Println("请输入矩阵mat：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&str1)
	mat, _ = utils.ParseMatrixString(str1)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入k：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&k)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘矩阵中战斗力最弱的 K 行’函数")
	result := kWeakestRows(mat, k)
	fmt.Println("结果为：")
	fmt.Println(result)
	utils.PrintName("‘矩阵中战斗力最弱的 K 行’程序结束")
}
