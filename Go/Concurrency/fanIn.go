package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration, s string) {
	var i int
	for {
		ch <- i
		fmt.Print(fmt.Sprintf("[%s - %d]", s,i))
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x:= range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	out := make (chan int)

	go producer(ch, 100*time.Millisecond, "a")
	go producer(ch, 250*time.Millisecond, "b")
	go reader(out)

	for i := range ch {
		out <- i
	}
}
