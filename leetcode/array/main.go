// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:38
// @Description: 数组类主函数主程序
package main

import (
	"fmt"
	"golang-learn/leetcode/array/function"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 20)
	fmt.Println("请输入数组值：")
	var input string
	fmt.Scanln(&input)

	values := strings.Split(input, ",")
	for i, v := range values {
		num, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Println("输入的值不是有效的整数:", v)
			return
		}
		nums[i] = num
	}

	// twoSum：输入数组和target，查询两数之和为target的角标
	var target int
	fmt.Println("开始‘两数之和’函数")
	fmt.Print("请输入target值：")
	fmt.Scanf("%d", &target)
	result := function.TwoSum(nums, target)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}

}
