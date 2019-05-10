package main

/*
	计算时间差

*/

import (
	"fmt"
	"time"
)

func main() {
	temp := 0
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		for j := 0; j < 10000; j++ {
			temp++
		}
	}
	d, _ := time.ParseDuration("-24h")
	fmt.Println(temp)
	t2 := time.Now()
	k := time.Now()
	fmt.Println(t2.Sub(t1))

	fmt.Println(k.Add(d)) //一天前
	test()                //调用一个函数的使用
	hh = 8
	fmt.Println(hh)

	fmt.Println("这是测试for test...")
	test_go()

}
