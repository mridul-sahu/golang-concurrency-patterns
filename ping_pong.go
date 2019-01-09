package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func PingPong() {
	table := make(chan *Ball)
	over := make(chan bool)
	go player("ping", table, over)
	go player("pong", table, over)

	table <- new(Ball)
	time.Sleep(time.Second)
	over <- true
	over <- true
}

func player(name string, table chan *Ball, over <-chan bool) {
	for {
		select {
		case ball := <-table:
			fmt.Println(name, ball.hits)
			ball.hits += 1
			time.Sleep(time.Millisecond * 100)
			table <- ball
		case <-over:
			return
		}
	}
}
