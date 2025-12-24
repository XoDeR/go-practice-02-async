package main

import "fmt"

func worker() int {
	return 345
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- worker()
	}()

	result := <-ch
	fmt.Println(result)
}
