package main

//同意一个包内是不可以有俩个main函数的
import "fmt"

var hh int //可以再main.go中使用

func test() {
	hh := "ee" //局部变量
	fmt.Println("e")
	fmt.Println(hh)

	fmt.Println("这已经可以结束了。")

	fmt.Println("这是 function testing ...")
	fmt.Println("hello world !!!!")

}
