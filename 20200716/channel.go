package main

import (
	"fmt"
	"time"
)

func main() {
	var channle chan int
	fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channle, channle) //

	// 使用 `make(chan type)` 创建一个新的通道。通道类型就是他们需要传递值的类型。
	message := make(chan string)

	// 使用 `channel <-` 语法 _发送(send)_ 一个新的值到通道中。这里
	// 我们在一个新的 Go 协程中发送 `"ping"` 到上面创建的
	// `messages` 通道中。
	go func() {
		message <- "ping"
	}()

	// 使用 `<-channel` 语法从通道中 _接收(receives)_ 一个值。这里
	// 将接收我们在上面发送的 `"ping"` 消息并打印出来。
	msg := <- message
	fmt.Println(msg)

	done := make(chan bool, 1)
	go worker(done)

	// 在接收到通道中 worker 发出的通知前一直阻塞。
	<- done
}

func worker(done chan bool)  {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}


