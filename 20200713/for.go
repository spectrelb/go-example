package main

import "fmt"

func main() {
	// 最基础的方式，单个循环条件。
	i := 1
	for i < 3  {
		fmt.Println(i)
		i += 1
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环将一直重复执行，遇到 `break` 或者 `return` 跳出循环。
	for  {
		fmt.Println("loop")
		break
	}

	for j := 1; j <= 3; j++ {
		if j % 2 != 0 {
			continue
		}
		fmt.Println(j)
	}
}
