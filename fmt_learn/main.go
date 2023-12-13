package main

import "golang-learn/fmt_learn/learn"

func main() {
	// 学习继承的使用
	// 父类作为子类的属性，子类可以直接调用父类的方法
	dog := learn.Dog{
		Animal: learn.Animal{Name: "dog", Age: 1},
	}
	dog.Eat()

	// 学习接口的实现
	var user learn.User
	// 这里我们看到"&"取址符号，它相当于Java的new。对象的创建存放在堆上，需要使用指针来指向
	user = &learn.Customer{}
	user.Save()
}
