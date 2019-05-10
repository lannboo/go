package main

import (
	"fmt"
	"time"
)

/*
select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。

select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。


*/

func chanTest(ch chan int) {

	for { //select外层需要循环
		select {
		case value, ok := <-ch:
			fmt.Println(value, ok, time.Now())
			if ok == false {
				fmt.Println("chan已经关闭", time.Now()) //select要自己判断退出，如果是for..range 形式，在读取完了关闭的chanel后，退出循环
				return
			}

		default:
			fmt.Println("chan 空了", time.Now())
			time.Sleep(time.Second * 5) //分支的处理会阻塞整个select
		}
	}
}

func main() {

	var ch = make(chan int, 100)

	go chanTest(ch)

	ch <- 1
	ch <- 2
	time.Sleep(time.Second * 2)
	ch <- 3
	ch <- 4

	time.Sleep(time.Second)

	close(ch)

	for {
		time.Sleep(time.Second)
	}
}
