package function

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	// 创建一个空节点
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	p := this // p指向虚拟头结点
	for i := 0; i <= index; i++ {
		if p != nil {
			p = p.Next
		} else {
			return -1
		}
	}
	if p != nil {
		return p.Val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	linkedList := &MyLinkedList{Val: val, Next: this.Next}
	this.Next = linkedList
}

func (this *MyLinkedList) AddAtTail(val int) {
	p := this
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	linkList := &MyLinkedList{Val: val}
	p := this
	for i := 0; i < index; i++ {
		if p != nil {
			p = p.Next
		} else {
			return
		}
	}
	if p != nil {
		linkList.Next = p.Next
		p.Next = linkList
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	p := this
	for i := 0; i < index; i++ {
		if p != nil {
			p = p.Next
		} else {
			return
		}
	}
	if p != nil && p.Next != nil {
		p.Next = p.Next.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
