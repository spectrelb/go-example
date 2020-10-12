package main

import (
	"fmt"
)

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

	b := 123
	fmt.Println(b)

	zeroval(b)
	fmt.Println("zeroval:", b)

	zeroptr(&b)
	fmt.Println("zeroval:", b)
}

// 有一个 `int` 型参数，所以使用值传递。将从调用它的那个函数中得到一个 `ival` 形参的拷贝。
func zeroval(ival int) {
	ival = 0
}

//`*int` 参数，意味着它用了一个 `int` 指针。`*iptr` 接着_解引用_这个指针，从它内存地址得到这个地址对应的当前值。
//对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}