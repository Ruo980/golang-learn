package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func detectCycle(head *utils.ListNode) *utils.ListNode {
	p1 := head
	p2 := head
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			p1 = head
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}
	return nil
}

func Exec142() {
	// 填充式输出题目名称
	title := "环形链表II"
	utils.PrintName(title)

	// 接收参数
	var str string
	var pos int
	fmt.Println("请输入链表：")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	fmt.Println("请输入pos：")
	_, err = fmt.Scanln(&pos)
	if err != nil {
		return
	}
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := detectCycle(head)
	// 打印结果
	if result != nil {
		fmt.Println(result.Val)
	} else {
		fmt.Println("nil")
	}
	//算法结束
	utils.PrintName("‘环形链表II’程序结束")
}
