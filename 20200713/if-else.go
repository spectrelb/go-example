package main

import "fmt"

func main() {
	// 这里是一个基本的例子。
	if 7%2 ==0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 你可以不要 `else` 只用 `if` 语句。
	if 8%4 ==0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句之前可以有一个声明语句；在这里声明的变量, 可以在所有的条件分支中使用。
	if a := 9; a < 10 {
		fmt.Println(a, "is negative")
	} else if a < 110 {
		fmt.Println(a, "has 1 digit")
	} else {
		fmt.Println(a, "has multiple digits")
	}
}
