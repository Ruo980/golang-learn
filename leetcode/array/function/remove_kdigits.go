package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := []byte{} //存放结果集
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// 结果集转为字符串
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func Exec402() {
	// 填充式输出题目名称
	title := "移掉K位数字"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var num string
	var k int
	fmt.Println("请输入字符串num：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&num)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入整数k：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&k)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘移掉K位数字’函数")
	result := removeKdigits(num, k)
	fmt.Println("结果为：")
	// 将字节byte转为字符输出
	fmt.Println(string(result))
	utils.PrintName("‘移掉K位数字’程序结束")
}
