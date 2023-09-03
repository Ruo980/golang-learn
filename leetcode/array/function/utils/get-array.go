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
