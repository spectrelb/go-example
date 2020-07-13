package main

import (
	"fmt"
	"math"
)

// const 用于声明一个常量
const a string = "constant"

func main() {
	fmt.Println(a)

	const  b = 123
	fmt.Println(b)

	//常数表达式可以执行任意精度的运算
	const  c = 100 / 3
	fmt.Println(c)

	//数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(float64(c))

	//函数调用方式获取pi值
	fmt.Println(math.Pi)

	//iota特殊的常量,iota是常量里面的计数器，初始值默认值是0，可以被编译器自动修改，每定义一组常量时，iota逐行自增1。
	const (
		A = iota  //为 0
		B       //默认和上一行一样，iota加1   为1
		C		//默认和上一行一样，iota加1   为2
	)

	const (
		D= iota
	)

	fmt.Println(A, B, C, D)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	fmt.Println(A, B, C, D)

	fmt.Println(1<<10)
}
