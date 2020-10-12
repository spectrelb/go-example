package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	name string
	age  int
}

type Students struct {
	ID     int 	  `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

type Class struct {
	Title    string        `json:"title"`
	Students []*Students   `json:"list"`
}

func main()  {
	//m := make(map[string]*student)
	//stus := []student{
	//	{name: "小王子", age: 18},
	//	{name: "娜扎", age: 23},
	//	{name: "大王八", age: 9000},
	//}
	//
	//for _, stu := range stus {
	//	m[stu.name] = &stu
	//}
	//
	///**
	//娜扎 => 大王八
	//大王八 => 大王八
	//小王子 => 大王八
	// */
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.name)
	//}

	c := &Class{
		Title:    "101",
		Students: make([]*Students, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Students{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"title":"101","list":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}


