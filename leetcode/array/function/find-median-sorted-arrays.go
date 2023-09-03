// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 21:23
// @Description: 寻找两个正序数组的中位数
// @input [1,2] [3,4]
package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 先合并两个数组再查找中位数 ：两个有序的数组，合并，时间复杂度为m+n，求的就是合并算法，在合并是能排好序
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	a := len(nums1)
	b := len(nums2)
	lens := a + b
	var value float64
	var nums []int
	for true {
		if i < a && j < b {
			if nums1[i] <= nums2[j] {
				nums = append(nums, nums1[i])
				i++
			} else {
				nums = append(nums, nums2[j])
				j++
			}
		} else {
			break
		}
	}
	for true {
		if i < a {
			nums = append(nums, nums1[i])
			i++
		} else {
			break
		}
	}
	for true {
		if j < b {
			nums = append(nums, nums2[j])
			j++
		} else {
			break
		}
	}
	if lens%2 == 0 {
		//nums偶数长
		value = float64((nums[lens/2]+nums[(lens-1)/2])*1.0) / 2
	} else {
		value = float64(nums[lens/2])
	}
	return value
}

// findMedianSortedArrays：合并两个数组找中间数
func Exec4() {
	// 填充式输出题目名称
	title := "Median of Two Sorted Arrays"
	utils.PrintName(title)

	//===开始输入相关参数
	fmt.Println("请输入数组1的值")
	var input1 string
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input1)
	reader.Discard(reader.Buffered())
	nums1 := utils.GetArray(input1)
	fmt.Println("请输入数组2的值")
	var input2 string
	fmt.Scanln(&input2)
	reader.Discard(reader.Buffered())
	nums2 := utils.GetArray(input2)

	// ==执行算法
	result2 := findMedianSortedArrays(nums1, nums2)

	// ==输出结果
	fmt.Println("结果如下：")
	fmt.Println(result2)
	utils.PrintName("“寻找两个正序数组的中位数”程序结束")
}
