package main

import (
	"fmt"
	"math"
)

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们就需要实现接口中的所有方法。 这里我们在 `rect` 上实现了 `geometry` 接口。
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64  {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64  {
	return 2 * math.Pi * c.radius
}

// 如果一个变量具有接口类型，那么我们可以调用指定接口中的方法。
func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{3, 4}
	c := circle{5}

	// 结构体类型 `circle` 和 `rect` 都实现了 `geometry` 接口，
	// 所以我们可以使用它们的实例作为 `measure` 的参数。
	measure(r)
	measure(c)
}
