package learn

import "fmt"

// Golang 继承的学习

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Eat() {
	fmt.Println("Animal Eat")
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}
