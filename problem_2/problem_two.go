package main

import "fmt"

func iterateSeq(ch chan int) {
	for {
		res, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
			break
		} else {
			fmt.Print(res, "\t")
		}
	}
	defer close(ch)
}
func main() {
	//fmt.Println("Start main method")
	//ch := make(chan int)
	//go myfunc(ch)
	//ch <- 79
	//fmt.Println("End main method")
	ch := make(chan int)
	values := []int{10, 20, 35, 100, 200, 502}
	go iterateSeq(ch)
	for i := 0; i < len(values); i++ {
		ch <- i
	}
}
