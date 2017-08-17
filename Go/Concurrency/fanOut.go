package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <- chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		task, ok := <- tasksCh
		if !ok { return; }
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Task: %d --  %d\n", task, i)
	}
}

func pool(wg *sync.WaitGroup, workers int, tasks int) {
	taskCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(taskCh, wg, i)
	}

	for i := 0; i < tasks; i++ {
		taskCh <- i
	}

	close(taskCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
	fmt.Println("END")
}
