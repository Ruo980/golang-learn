// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-01 22:37
// @Description:
package main

import "fmt"

func main() {
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}
