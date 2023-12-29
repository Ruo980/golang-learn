package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	count := 0      // 计数器
	j := len(s) - 1 // s的下标

	// 排序
	sort.Ints(g)
	sort.Ints(s)

	for i := len(g) - 1; i >= 0; i-- {
		// 先判断饼干是不是发完了
		if j == -1 {
			return count
		}
		// 如果饼干大于等于胃口值，计数器加一，饼干下标减一
		if g[i] <= s[j] {
			count++
			j--
		}
	}
	return count
}

func Exec455() {
	// 填充式输出题目名称
	title := "分发饼干"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var g []int
	var s []int
	var str1 string
	var str2 string
	fmt.Println("请输入孩子的胃口值：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&str1)
	g = utils.GetArray(str1)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("请输入饼干的尺寸：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&str2)
	s = utils.GetArray(str2)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘分发饼干’函数")
	result := findContentChildren(g, s)
	fmt.Println("结果为：")
	fmt.Println(result)
	utils.PrintName("‘分发饼干’程序结束")
}
