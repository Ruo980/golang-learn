package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	// 创建一个虚拟头节点
	p1 := &utils.ListNode{}
	p1.Next = head
	p2 := p1
	for head != nil {
		if head.Val == val {
			p2.Next = head.Next
		} else {
			p2 = p2.Next
		}
		head = head.Next
	}
	return p1.Next
}

func Exec203() {
	// 填充式输出题目名称
	title := "移除链表元素"
	utils.PrintName(title)

	// 接收参数
	var str string
	var val int
	fmt.Println("请输入链表：")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	fmt.Println("请输入要删除的元素：")
	_, err = fmt.Scanln(&val)
	if err != nil {
		return
	}
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := removeElements(head, val)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘移除链表元素’程序结束")
}
