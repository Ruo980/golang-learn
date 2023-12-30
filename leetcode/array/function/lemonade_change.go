package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func lemonadeChange(bills []int) bool {
	a := 0
	b := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			a++
		} else if bills[i] == 10 {
			// 找5
			a = a - 1
			b++
			if a < 0 {
				return false
			}
		} else if bills[i] == 20 {
			// 找15：10+5或者5*3
			if a > 0 && b > 0 {
				a--
				b--
			} else if a >= 3 {
				a = a - 3
			} else {
				return false
			}
		}
	}
	return true
}

func Exec860() {
	// 填充式输出题目名称
	title := "柠檬水找零"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作：输入数组
	fmt.Println("请输入数组值：")
	var input string
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input)
	// 输入格式为：[-2,1,-3,4,-1,2,1,-5,4]
	nums := utils.GetArray(input)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘柠檬水找零’函数")
	result := lemonadeChange(nums)
	fmt.Println("结果为：", result)
	utils.PrintName("‘柠檬水找零’程序结束")
}
