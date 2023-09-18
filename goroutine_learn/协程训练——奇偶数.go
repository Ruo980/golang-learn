// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-01 19:00
// @Description: 协程相关训练，利用channel让两个协程协作输出0~100奇偶数
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道
	ch := make(chan int)
	// 偶数输出数量
	a := 0
	// 奇数输出数量
	b := 0
	// 输出偶数的协程
	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				fmt.Println("协程1输出：", i)
				// 一直阻塞等待数据传输：当阻塞超时抛出异常，当顺畅时继续进行
				ch <- i
				a++
			}
		}
		close(ch)
	}()

	// 输出奇数的协程
	go func() {
		for i := 0; i < 100; i++ {
			// 建立连接：接收传输值
			<-ch
			if i%2 != 0 {
				fmt.Println("协程2输出奇数：", i)
				b++
			}
		}
	}()

	time.Sleep(3 * time.Second)
	// 判断输出个数是否正常
	fmt.Println(a, b)
}
