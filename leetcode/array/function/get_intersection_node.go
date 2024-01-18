package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	p1 := headA
	p2 := headB
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		p1 = p1.Next
		headA = headA.Next
	}
	for p2 != nil {
		p2 = p2.Next
		headB = headB.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func Exec0207() {
	// 填充式输出题目名称
	title := "相交链表"
	utils.PrintName(title)

	// 接收参数
	var str1, str2 string
	var pos int
	fmt.Println("请输入链表1：")
	_, err := fmt.Scanln(&str1)
	if err != nil {
		return
	}
	fmt.Println("请输入链表2：")
	_, err = fmt.Scanln(&str2)
	if err != nil {
		return
	}
	fmt.Println("请输入pos：")
	_, err = fmt.Scanln(&pos)
	if err != nil {
		return
	}
	pos = 2
	//转换为链表
	headA := utils.GetList(str1)
	headB := utils.GetList(str2)
	// 执行算法
	result := getIntersectionNode(headA, headB)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘相交链表’程序结束")
}
