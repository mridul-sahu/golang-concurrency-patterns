package main

import (
	"fmt"
	"time"
)

type Request struct {
	query string
	data chan string
}

func SendWhenAvailable() chan<-Request {
	r := make(chan Request)

	go func(r <-chan Request) {
		for {
			request := <-r
			time.Sleep(time.Millisecond * 100)
			request.data <- fmt.Sprintf("%s is a very intresting question indeed!", request.query)
		}
	}(r)

	return r
}
