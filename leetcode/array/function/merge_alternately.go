package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	if len(word1) > len(word2) {
		for i := 0; i < len(word2); i++ {
			result += string(word1[i]) + string(word2[i])
		}
		result += word1[len(word2):]
	} else {
		for i := 0; i < len(word1); i++ {
			result += string(word1[i]) + string(word2[i])
		}
		result += word2[len(word1):]
	}
	return result

}

func Exec1768() {
	// 填充式输出题目名称
	title := "交替合并字符串"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var word1 string
	var word2 string
	fmt.Println("请输入word1：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&word1)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入word2：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&word2)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘交替合并字符串’函数")
	result := mergeAlternately(word1, word2)
	fmt.Println("结果为：")
	fmt.Println(result)
	utils.PrintName("‘交替合并字符串’程序结束")
}
