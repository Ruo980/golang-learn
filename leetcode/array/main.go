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

func getArray(input string) []int {
	fmt.Println(input)
	nums := make([]int, 20)
	values := strings.Split(input, ",")
	for i, v := range values {
		num, _ := strconv.Atoi(strings.TrimSpace(v))
		nums[i] = num
	}
	fmt.Println("dede")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		fmt.Println("a")
	}
	return nums
}

func main() {

	// twoSum：输入数组和target，查询两数之和为target的角标
	/*	fmt.Println("请输入数组值：")
		var input string
		fmt.Scanln(&input)
		nums := getArray(input)
		fmt.Println("开始‘两数之和’函数")
		fmt.Print("请输入target值：")
		var target int
		fmt.Scanf("%d", &target)
		fmt.Println(target)

		result := function.TwoSum(nums, target)
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i])
			fmt.Print(" ")
		}*/

	fmt.Println()
	// findMedianSortedArrays：合并两个数组找中间数
	fmt.Println("开始‘两有序数组找中间数’函数")
	fmt.Println("请输入数组1的值")
	var input1 string
	fmt.Scanln(&input1)
	nums1 := getArray(input1)

	fmt.Println("请输入数组2的值")
	var input2 string
	fmt.Scanln(&input2)
	nums2 := getArray(input2)
	result2 := function.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(result2)
}
