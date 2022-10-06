package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, myChan chan bool) {
	myChan <- true
	x = x + 1
	<-myChan
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	myChan := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, myChan)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
