package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func runningSum(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}

func Exec1480() {
	// 填充式输出题目名称
	title := "Dynamic sum of one-dimensional arrays"
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
	fmt.Println("开始‘一维数组的动态和’函数")
	max := runningSum(nums)
	fmt.Println("一维数组的动态和：")
	fmt.Println(max)
	utils.PrintName("‘一维数组的动态和’程序结束")
}
