package main

import "fmt"

func main() {
	a := make(chan int, 5)
	b := make(chan bool)

	go func() {
		for  {
			i, ok := <- a
			if ok {
				fmt.Println("received job", i)
			} else {
				fmt.Println("received all jobs")
				b <- true
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		a <- i
		fmt.Println("send", i)
	}

	close(a)
	fmt.Println("send all")

	<- b
}
