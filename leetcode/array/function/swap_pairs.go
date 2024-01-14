package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func swapPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l1 := head.Next
	l2 := head.Next
	l3 := head.Next
	head.Next = nil
	for l1 != nil {
		l1 = l1.Next
		l2.Next = head
		if l1 == nil || l1.Next == nil {
			head.Next = l1
			break
		} else {
			head.Next = l1.Next
			head = l1
			l1 = l1.Next
			l2 = l1
		}
	}
	return l3
}

func Exec24() {
	// 填充式输出题目名称
	title := "两两交换链表中的节点"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入链表：")
	fmt.Scanln(&str)
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := swapPairs(head)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘两两交换链表中的节点’程序结束")
}
