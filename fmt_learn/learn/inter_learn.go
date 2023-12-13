package learn

import "fmt"

// Golang接口的实现
// 定义一个接口
type User interface {
	Save()
}

type Customer struct {
}

func (c *Customer) Save() {
	fmt.Println("customer save")
}
