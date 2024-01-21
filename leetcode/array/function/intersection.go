package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func intersection(nums1 []int, nums2 []int) []int {
	// 创建一个map集合
	m := make(map[int]int, len(nums1))
	// 遍历nums1，将元素存入map集合
	for _, v := range nums1 {
		m[v]++
	}
	// 创建一个切片，用于存放结果
	var res []int
	// 遍历nums2，判断元素是否在map集合中
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v] = 0 // 必须清空，放置重复元素进入
		}
	}
	return res
}

func Exec349() {
	// 填充式输出题目名称
	title := "两个数组的交集"
	utils.PrintName(title)

	// 接收参数
	var str1, str2 string
	fmt.Println("请输入数组nums1：")
	_, err := fmt.Scanln(&str1)
	if err != nil {
		return
	}
	fmt.Println("请输入数组nums2：")
	_, err = fmt.Scanln(&str2)
	if err != nil {
		return
	}
	//转换为数组
	nums1 := utils.GetArray(str1)
	nums2 := utils.GetArray(str2)
	// 执行算法
	result := intersection(nums1, nums2)
	// 打印结果
	fmt.Println("两个数组的交集为：", result)
	//算法结束
	utils.PrintName("‘两个数组的交集’程序结束")
}
