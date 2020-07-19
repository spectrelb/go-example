package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(time.Second)
		a <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		b <- "two"
	}()

	fmt.Println("结束。。。")

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <- a:
			fmt.Println(msg1)
		case msg2 := <- b:
			fmt.Println(msg2)
			// `<-Time.After` 等待超时, 时间 2 秒后发送的值。
		case <- time.After(time.Second * 2):
			fmt.Println("已经超时")
		default:
			fmt.Println("完成。。。。。")
		}
	}
}


