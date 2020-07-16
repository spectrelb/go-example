package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//获取当前GOROOT目录
	fmt.Println("GOROOT:", runtime.GOROOT())
	//获取当前操作系统
	fmt.Println("操作系统:", runtime.GOOS)
	fmt.Println("逻辑CPU数量：", runtime.NumCPU())
	go testgo1()
	go testgo2()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main 函数结束")
}

func testgo1 ()  {
	for i := 0; i < 10; i++ {
		fmt.Println("测试goroutine1",i)
	}
}

func testgo2 ()  {
	for i := 0; i < 10; i++ {
		fmt.Println("测试goroutine2",i)
	}
}