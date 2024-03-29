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
	fmt.Println("欢迎进入 leetcode 刷题：")
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
		case 19:
			function.Exec19() // 执行题目编号为 19 的程序
		case 24:
			function.Exec24() // 执行题目编号为 24 的程序
		case 26:
			function.Exec26() // 执行题目编号为 26 的程序
		case 80:
			function.Exec80() // 执行题目编号为 80 的程序
		case 27:
			function.Exec27() // 执行题目编号为 27 的程序
		case 53:
			function.Exec53() // 执行题目编号为 53 的程序
		case 83:
			function.Exec83() // 执行题目编号为 83 的程序
		case 0207:
			function.Exec0207() // 执行题目编号为 0207 的程序
		case 142:
			function.Exec142() // 执行题目编号为 142 的程序
		case 151:
			function.Exec151() // 执行题目编号为 151 的程序
		case 203:
			function.Exec203() // 执行题目编号为 203 的程序
		case 206:
			function.Exec206() // 执行题目编号为 206 的程序
		case 209:
			function.Exec209() // 执行题目编号为 209 的程序
		case 242:
			function.Exec242() // 执行题目编号为 242 的程序
		case 344:
			function.Exec344() // 执行题目编号为 344 的程序
		case 349:
			function.Exec349() // 执行题目编号为 349 的程序
		case 389:
			function.Exec389() // 执行题目编号为 389 的程序
		case 402:
			function.Exec402() // 执行题目编号为 402 的程序
		case 455:
			function.Exec455() // 执行题目编号为 455 的程序
		case 674:
			function.Exec674() // 执行题目编号为 674 的程序
		case 860:
			function.Exec860() // 执行题目编号为 860 的程序
		case 976:
			function.Exec976() // 执行题目编号为 976 的程序
		case 977:
			function.Exec977() // 执行题目编号为 977 的程序
		case 1143:
			function.Exec1143() // 执行题目编号为 95 的程序
		case 1209:
			function.Exec1209() // 执行题目编号为 1209 的程序
		case 1337:
			function.Exec1337() // 执行题目编号为 1337 的程序
		case 1342:
			function.Exec1342() // 执行题目编号为 1342 的程序
		case 1480:
			function.Exec1480() // 执行题目编号为 1480 的程序
		case 1768:
			function.Exec1768() // 执行题目编号为 1768 的程序
		case 1844:
			function.Exec1844() // 执行题目编号为 1844 的程序
		default:
			utils.PrintName("未找到对应的题号，请重新输入！")
		}

	}
}
