package main

import "fmt"

func main() {

	// 这里我们创建了一个长度为5的 `int` 类型 slice（初始化为零值）。
	a := make([]int, 5)
	fmt.Println("emp:", a)

	// 我们可以和数组一样设置和得到值
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	// 使用内置函数 `len` 返回slice的长度。
	fmt.Println("len:", len(a))

	// 这是一个内建的 `append`，它返回一个包含了一个或者多个新值的 slice。 由于 `append` 可能返回新的 slice，我们需要接受其返回值。
	a = append(a, 6)
	a = append(a, 7,8,9)
	fmt.Println("emp:", a)

	// Slice 也可以被 `copy`。相同长度的 slice `b`，并且将 `a` 复制给 `b`。
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("copy:", b)

	// Slice 通过 `slice[low:high]` 语法进行“切片”操作。
	// 这个 slice 从 `a[0]` 切片到 `a[2]`（不包含）。
	c := a[0:2]
	fmt.Println("sli1:", c)

	// 这个 slice 从 `a[0]` 切片到 `a[2]`（不包含）。
	d := a[:2]
	fmt.Println("sli2:", d)

	// 这个 slice 从 `a[2]` 切片（包含）开始切片。
	e := a[2:]
	fmt.Println("sli3", e)

	// 使用短语法在一行内声明并初始化一个slice 。
	f := []string{"a", "b"}
	fmt.Println(f)

	//构建二维slice,每一个slice的长度可以不一致
	two := [][]string{{"a"}, {"b", "c"},{"d", "e","f"}}
	fmt.Println(two)
}
