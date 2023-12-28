// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:38
// @Description: 数组类主函数主程序
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang-learn/leetcode/array/function"
	"golang-learn/leetcode/array/function/utils"
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
			function.Exec11() // 执行题目编号为 11 的程序
		case 15:
			function.Exec15() // 执行题目编号为 15 的程序
		case 16:
			function.Exec16() // 执行题目编号为 16 的程序
		case 53:
			function.Exec53() // 执行题目编号为 53 的程序
		case 1768:
			function.Exec1768() // 执行题目编号为 1768 的程序
		case 389:
			function.Exec389() // 执行题目编号为 389 的程序
		case 674:
			function.Exec674() // 执行题目编号为 674 的程序
		case 1480:
			function.Exec1480() // 执行题目编号为 1480 的程序
		case 1342:
			function.Exec1342() // 执行题目编号为 1342 的程序
		default:
			utils.PrintName("未找到对应的题号，请重新输入！")
		}

	}
}
