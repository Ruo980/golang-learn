package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func findTheDifference(s string, t string) byte {
	var res byte
	var a = [1001]int{0}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if a[j] != 1 && s[i] == t[j] {
				a[j] = 1
				break
			}
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			res = t[i]
			break
		}
	}
	return res
}

func Exec389() {
	// 填充式输出题目名称
	title := "找不同"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var s string
	var t string
	fmt.Println("请输入字符串s：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&s)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入字符串t：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&t)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘找不同’函数")
	result := findTheDifference(s, t)
	fmt.Println("结果为：")
	// 将字节byte转为字符输出
	fmt.Println(string(result))
	utils.PrintName("‘交替合并字符串’程序结束")
}
