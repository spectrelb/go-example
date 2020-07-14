package main

import "fmt"

func main() {
	// range可以用来便利数组，切片，map等数据结构
	a := []int{2, 3, 4}
	for k, v := range a {
		fmt.Printf("%v -> %v\n", k, v)
	}

	// range遍历，可以忽略key
	for _, v := range a {
		fmt.Printf("%v\n", v)
	}

	// `range` 在 map 中迭代键值对。
	b := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range b {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
