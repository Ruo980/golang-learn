package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func removeDuplicates3(s string, k int) string {
	// 字符串转换为数组：表示栈和待插入序列
	str := []byte(s)

	nums := make([]int, len(str))
	t := -1 // 当前栈顶的长度

	// 起始情况下：栈顶元素为str[0],长度为1,待插入元素为str[1]
	i := -1 // i为慢指针,表示栈顶
	j := 0  // j为快指针,表示当前待插入的元素

	// 开始插入元素
	for j < len(str) {
		// 判断当前元素是否与栈顶元素相同
		if i != -1 && str[i] == str[j] {
			// 判断当前的长度是否超过k-1:如果超过k-1，则将栈顶k-1个元素移除;否则将栈顶元素的长度加1
			if nums[t] == k-1 {
				i = i - k + 1
				t-- // 上一个不同与栈顶元素的元素积累的长度
			} else {
				i++
				str[i] = str[j]
				nums[t]++
			}
		} else {
			// 如果当前元素与栈顶元素不同，则将当前元素加入到栈中
			i++
			str[i] = str[j]
			t++
			nums[t] = 1
		}
		j++
	}
	return string(str[:i+1])
}

func Exec1209() {
	// 填充式输出题目名称
	title := "删除字符串中的所有相邻重复项 II"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var s string
	var k int
	fmt.Println("请输入字符串s：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	_, err := reader.Discard(reader.Buffered())
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&s)
	if err != nil {
		return
	}
	// 清空输入缓冲区
	_, err = reader.Discard(reader.Buffered())
	if err != nil {
		return
	}
	fmt.Println("请输入整数k：")
	_, err = fmt.Scanln(&k)
	if err != nil {
		return
	}
	fmt.Println("开始‘删除字符串中的所有相邻重复项 II’函数")
	str := removeDuplicates3(s, k)
	fmt.Println("删除字符串中的所有相邻重复项 II后的字符串为：", str)
	utils.PrintName("‘删除字符串中的所有相邻重复项 II’程序结束")
}
