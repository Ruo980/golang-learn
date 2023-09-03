// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:41
// @Description: 两数之和
// @input [1,] [3,4]
package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

// 给定一个数组和target。在数组中找到两个数，它们之和为target。这个数组中符合条件的两数对仅一个
func twoSum(nums []int, target int) []int {
	// 利用map 离散存放值的特点
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			// 找到对应存放的another返回
			return []int{i, m[another]}
		} else {
			//如果没找到对应的值，就把现在的值存起来
			m[nums[i]] = i
		}
	}

	// 遍历结束仍未找到，返回nil
	return nil
}

// twoSum：输入数组和target，查询两数之和为target的角标
func Exec1() {
	// 填充式输出题目名称
	title := "Two Sum"
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
	fmt.Println("开始‘两数之和’函数")
	fmt.Print("请输入target值：")
	var target int
	fmt.Scanf("%d", &target)

	// ==执行算法
	result := twoSum(nums, target)

	// ==输出结果
	fmt.Println("结果如下：")
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}
	fmt.Println()
	utils.PrintName("”两数之和“程序结束")
}
