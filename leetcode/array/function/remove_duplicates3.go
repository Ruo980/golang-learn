package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func removeDuplicates3(s string, k int) string {
	var strStack []byte
	var countStack []int
	strStack = append(strStack, s[0])
	countStack = append(countStack, 1)

	for i := 1; i < len(s); i++ {
		if len(strStack) > 0 && s[i] == strStack[len(strStack)-1] {
			countStack[len(countStack)-1]++
		} else {
			countStack = append(countStack, 1)
		}
		strStack = append(strStack, s[i])

		if countStack[len(countStack)-1] == k {
			strStack = strStack[0 : len(strStack)-k]
			countStack = countStack[0 : len(countStack)-1]
		}
	}
	return string(strStack)
}

func Exec1209() {
	// 填充式输出题目名称
	title := "删除字符串中的所有相邻重复项 II"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var s string
	var k int
	fmt.Println("请输入字符串s：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&s)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入整数k：")
	fmt.Scanln(&k)
	fmt.Println("开始‘删除字符串中的所有相邻重复项 II’函数")
	str := removeDuplicates3(s, k)
	fmt.Println("删除字符串中的所有相邻重复项 II后的字符串为：", str)
	utils.PrintName("‘删除字符串中的所有相邻重复项 II’程序结束")
}
