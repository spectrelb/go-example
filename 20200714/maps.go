package main

import "fmt"

func main() {

	// 要创建一个空 map, `make(map[key-type]val-type)`.
	a := make(map[int]int)
	fmt.Println(a)

	// 使用典型的 `make[key] = val` 语法来设置键值对。
	a[1] = 1
	a[2] = 2
	fmt.Println("map:", a)

	// 我们可以和数组一样设置和得到值
	a[1] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	// 使用内置函数 `len` 返回map的长度。
	fmt.Println("len:", len(a))

	// 内建的 `delete` 可以从一个 map 中移除键值对
	delete(a, 1)
	fmt.Printf("map: %v, len:%v\n", a, len(a))

	// 使用短语法在一行内声明并初始化一个map。
	b := map[int]int{1:1, 2:2}
	fmt.Println(b)
}
