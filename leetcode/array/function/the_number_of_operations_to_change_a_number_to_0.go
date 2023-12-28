package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		count += 1
	}
	return count
}

func Exec1342() {
	// 填充式输出题目名称
	title := "Number of steps to change a number to 0"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作：输入数组
	fmt.Println("请输入数值：")
	var input int
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input)
	// 输入格式为：[-2,1,-3,4,-1,2,1,-5,4]
	fmt.Println("开始‘将数字变成0的操作次数’函数")
	count := numberOfSteps(input)
	fmt.Println("操作次数为：", count)
	utils.PrintName("‘将数字变成0的操作次数’程序结束")

}
