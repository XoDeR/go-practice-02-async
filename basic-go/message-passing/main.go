package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	v := <-ch
	fmt.Println(v)
}
