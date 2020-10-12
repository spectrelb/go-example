package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("current time:%v", now)
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(now.Unix(), 0).Format("2006-01-02 15:04:05"))

	fmt.Println(now.Add(time.Hour))
	fmt.Println(now.Add(time.Hour * 24))

	now1 := time.Date(2020, time.April, 10, 9,10,11, 0, time.UTC)
	fmt.Println(now.After(now1))
	fmt.Println(now.Before(now1))
	fmt.Println(now.Sub(now1))
}
