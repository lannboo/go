package main

import "fmt"

func main() {
	var a int = 3

	fmt.Println("hello world@@@")
	if a == 3 {
		fmt.Println(a)
	} else {
		fmt.Println("a is not is 3")
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	s := "adcdefg"
	for i := range s {
		fmt.Printf("%c", s[i])
	}
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	for i, c := range s {
		fmt.Printf("%d%c\n", i, c)
	}

	for i := 0; i < 5; i++ {
		if 2 == i {
			break //跳出循环
			//continue  //跳出本次循环
		}
		fmt.Println(i)
	}
	fmt.Println("goto LABEL --------------\n")
	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			goto LABEL //跳转到标签LABEL，从标签处，执行代码     直接跳出
		}
	}

	fmt.Println("this is test")

LABEL:
	fmt.Println("it is over")

}
