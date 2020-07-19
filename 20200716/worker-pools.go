package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i < 4; i++ {
		go workerPool(i, jobs, results)
	}

	for j := 1; j < 9; j++ {
		jobs <- j
	}
	close(jobs)

	for m := 1; m < 10; m++ {
		fmt.Println(<- results)
	}
}

func workerPool(id int, jobs <- chan int, results chan<- int)  {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j*2
	}
}