// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-03 11:47
// @Description: 输入数组字符串形式：[1,2,3]，转换为数组变量
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GetArray 将字符串解析为一维数组：传入格式为[1,2,3,4]
func GetArray(input string) []int {
	fmt.Println("输入input为" + input)
	input = strings.Trim(input, "[]")   // 移除字符串前后的方括号
	values := strings.Split(input, ",") // 将字符串按逗号分割为字符串列表
	nums := make([]int, 0, len(values)) // 创建一个大小为字符串列表长度的整数数组
	for _, v := range values {
		num, _ := strconv.Atoi(strings.TrimSpace(v))
		nums = append(nums, num) // 切片不能使用索引赋值的方式扩充空间
	}
	return nums
}

// GetArrayByString 将字符串解析为一维数组：传入格式为["h","e","l","l","o"]
func GetArrayByString(input string) []byte {
	fmt.Println("输入input为" + input)
	input = strings.Trim(input, "[]")   // 移除字符串前后的方括号
	values := strings.Split(input, ",") // 将字符串按逗号分割为字符串列表
	// 将得到的字符串转换为字符列表 []byte
	var nums []byte
	for _, v := range values {
		nums = append(nums, v[1])
	}
	fmt.Println("经过处理后的输入：", nums)
	return nums
}

// ParseMatrixString 解析字符串为二维整数切片:传入格式为[[1,2,3],[4,5,6],[7,8,9]]
func ParseMatrixString(input string) ([][]int, error) {
	// 去掉首尾的空格和换行符，并去掉方括号
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "[]")

	// 用逗号分割成元素
	elems := strings.Split(input, "],[")

	// 初始化二维切片
	var matrix [][]int

	// 遍历每个元素，将字符串转换成整数
	for _, rowStr := range elems {
		rowStr = strings.TrimSpace(rowStr)

		// 初始化一行
		var row []int

		// 遍历每个元素，将字符串转换成整数
		elemStrs := strings.Split(rowStr, ",")
		for _, elemStr := range elemStrs {
			elemStr = strings.TrimSpace(elemStr)
			num, err := strconv.Atoi(elemStr)
			if err != nil {
				return nil, fmt.Errorf("解析输入失败：%v", err)
			}
			row = append(row, num)
		}

		// 将这一行加入到二维切片中
		matrix = append(matrix, row)
	}

	return matrix, nil
}
