package main

import "fmt"

func main() {
	// 这里我们创建了一个数组 `a` 来存放刚好 5 个 `int`。 数组默认是零值的，对于 `int` 数组来说也就是 `0`。
	var  a [5]int
	fmt.Println("emp:", a)

	// 我们可以使用 `array[index] = value` 语法来设置数组，或者用 `array[index]` 得到值。
	a[1] = 222
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	// 使用内置函数 `len` 返回数组的长度。
	fmt.Println("len:", len(a))

	// 使用短语法在一行内声明并初始化一个数组。
	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	//构建二维数组
	two := [2][3]int{{1,2,3},{2,3,4}}
	fmt.Println(two)
}
