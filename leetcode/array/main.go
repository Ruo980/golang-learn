// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:38
// @Description: 数组类主函数主程序
package main

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"strconv"

	"golang-learn/leetcode/array/function"
)

func main() {
	fmt.Println("欢迎进入leetcode 数组类刷题：")
	for {
		fmt.Println("请输入题号进行相关程序演示 (按0退出)：")
		// 输入题号
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		questionNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("输入的题号无效，请重新输入！")
			continue
		}

		if questionNumber == 0 {
			break
		}

		switch questionNumber {
		case 1:
			function.Exec1() // 执行题目编号为 1 的程序
		case 4:
			function.Exec4() // 执行题目编号为 4 的程序
		case 11:
			function.Exec11() //执行题目编号为 11 的程序
		default:
			utils.PrintName("未找到对应的题号，请重新输入！")
		}

	}
}
