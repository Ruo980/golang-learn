package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func MaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	D := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		D[i] = make([]int, len(text2)+1)
		D[i][0] = 0
	}
	for i := 0; i < len(text2); i++ {
		D[0][i] = 0
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				D[i+1][j+1] = 1 + D[i][j]
			} else {
				D[i+1][j+1] = MaxValue(D[i][j+1], D[i+1][j])
			}
		}
	}
	return D[len(text1)][len(text2)]
}

func Exec1143() {
	// 填充式输出题目名称
	title := "longestCommonSubsequence"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var s string
	var t string
	fmt.Println("请输入字符串text1：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&s)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入字符串text2：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&t)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘最长公共子序列’函数")
	result := longestCommonSubsequence(s, t)
	fmt.Println("结果为：")
	// 将字节byte转为字符输出
	fmt.Println(result)
	utils.PrintName("‘最长公共子序列’程序结束")
}
