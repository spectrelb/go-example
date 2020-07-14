package main

import "fmt"

func main() {
	//定义一个变量
	a := 2
	fmt.Printf("变量A的地址为%p", &a) //通过%p占位符, &符号获取变量的内存地址。

	//创建一个指针，指针的声明 通过 *T 表示T类型的指针
	var i *int     //int类型的指针
	fmt.Println(i)
}
