package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch <- "1111"
	fmt.Println("added")
	a := <-ch
	fmt.Printf("%v\n", a)
}

/*
func main() {
	wg := &sync.WaitGroup{}
	cc := channel.NewChanCache()

	wg.Add(1)

	for i := 0; i < 10; i++ {
		ch := cc.ChannelGet("a")
		l := len(ch)
		fmt.Printf("Goroutines running %v, %v\n", runtime.NumGoroutine(), l)
	}
	time.Sleep(time.Second)
	fmt.Printf("after sleep Goroutines running %v\n", runtime.NumGoroutine())
}*/
