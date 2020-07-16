package main

import (
	json "encoding/json"
	"fmt"
)

// 这里的 `person` 结构体包含了 `name` 和 `age` 两个字段。
type person struct {
	name string
	age int
}

type Prescription struct {
	Name string
	Unit string
	Additive Prescription2
}

type Prescription2 struct {
	Name string
	Age int
}

type Prescription3 struct {
	Name string `json:"name"`
	Age int		`json:"age"`
}

func main() {

	// 省略的字段将被初始化为零值。
	fmt.Println(person{})

	//使用简短声明方式，后面加上{}代表这是结构体,直接给结构体成员变量赋值。
	p := person{"bobo", 28}
	fmt.Println(p)

	//使用点.来访问结构体内成员的变量的值。
	fmt.Println(p.name,p.age)

	//给结构体内成员变量赋值
	p1 := person{}
	p1.name = "bobo1"
	p1.age = 28
	fmt.Println(p1)

	//使用结构体指针
	p2 := person{name: "bobo2", age: 28}
	var pe *person
	pe = &p2
	fmt.Println(pe)

	//定义函数初始化结构体
	p3 := newPerson("bobo3", 11)
	fmt.Println(p3.name, p3.age)

	//嵌套结构体
	pre := Prescription{}
	pre.Name = "测试"
	pre.Unit = "测试"
	pre.Additive = Prescription2{"测试1", 18}
	fmt.Println(pre)

	//结构体与json转换
	//FIXME 需要注意的是将结构体转换为Json数据时候，定义结构体的字段必须首字母大写。否则无法正常解析。
	buf, err := json.Marshal(pre)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("json = ", string(buf))

	//可以看出其中json字符中每一个key的首字母也是大写，最后一个没有设置数据的字段的结果为null。那么如何强制将他变为小写的。并且将不需要显示的字段隐藏掉。就需要在结构体上添加标记。
	pre1 := Prescription3{"测试", 11}
	buf1, err1 := json.Marshal(pre1)
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}

	fmt.Println("json = ", string(buf1))

	//json与结构体转换
	jsonstr := `{"name":"测试","age":11}`
	var pre3 Prescription3
	if err2 := json.Unmarshal([]byte(jsonstr), &pre3); err2 !=nil {
		fmt.Println()
	}
	fmt.Println(pre3)
}

//返回结构体指针
func newPerson(name string, age int) *person  {
	return &person{
		name: name,
		age: age, //最后一个逗号不能省略
	}
}