package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func threeSumClosest(nums []int, target int) int {

	return 0
}

func Exec16() {
	// 填充式输出题目名称
	title := "3Sum Closest"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	fmt.Println("请输入数组值：")
	var input string
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input)
	nums := utils.GetArray(input)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘最接近的三数之和’函数")
	fmt.Print("请输入target值：")
	var target int
	fmt.Scanf("%d", &target)
	// 清空输入缓冲区
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// ==执行算法
	result := threeSumClosest(nums, target)
	fmt.Println(result)
	fmt.Println()
	utils.PrintName("”最接近的三数之和“程序结束")
}
