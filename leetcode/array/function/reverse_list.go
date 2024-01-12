package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

// reverseList
//
//	@Description: 反转链表
//	@param head
//	@return *ListNode
func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	l1 := head
	l2 := head.Next
	l1.Next = nil
	for l2 != nil {
		l1 = l2
		l2 = l2.Next
		l1.Next = head
		head = l1
	}
	return l1
}

func Exec206() {
	// 填充式输出题目名称
	title := "反转链表"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入链表：")
	fmt.Scanln(&str)
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := reverseList(head)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘反转链表’程序结束")
}
