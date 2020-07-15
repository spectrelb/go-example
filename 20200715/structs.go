package main

import "fmt"

func main() {
	a := 2
	fmt.Printf("变量A的地址为%p\n", &a) //通过%p占位符, &符号获取变量的内存地址。

	//创建一个指针，指针的声明 通过 *T 表示T类型的指针
	var i *int     //int类型的指针
	i = &a
	fmt.Println(i)  //i存储a的内存地址0xcXXXX
	fmt.Println(*i) //i存储这个指针存储的变量的数值2

	*i = 222
	fmt.Println(*i) //222
	fmt.Println(a) //222 通过指针操作 直接操作的是指针所对应的数值
}
