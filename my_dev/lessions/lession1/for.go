// for.go
package main

import (
	"fmt"
)

func test_go() {
	//顺序
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//
	for j := 0; j < 5; j++ {
		if j == 2 {
			break //跳出一个循环
			//continue   //跳出本次的一个循环，下一次会继续执行的
		}
		fmt.Println("j=", j)
	}

	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			goto LABEL //跳转到标签LABEL，从标签处，执行代码
		}
	}

	fmt.Println("this is test")

LABEL:
	fmt.Println("it is over")
	fmt.Print("jd\n")

}
