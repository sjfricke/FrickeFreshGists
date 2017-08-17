package main

import (
	"time"
	"log"
)

var a int

func player(table chan int, a int) {
	log.Println(a)
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		log.Printf("ball %d - [%d]", ball, a)
		table <- ball
	}
}

func main() {
	var Ball int
	table := make(chan int)
	go player(table, 1)
	go player(table, 2)
	go player(table, 3)

	table <- Ball
	time.Sleep(1 * time.Second)
	log.Println(<- table)
}
