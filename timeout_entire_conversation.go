package main

import (
	"fmt"
	"time"
)

func BoringTimeoutEntireConversation(input <-chan string) {
	timeout := time.After(time.Second * 10)
	for {
		select {
		case s := <-input:
			fmt.Print(s)
		case <-timeout:
			fmt.Print("That's it, I am done")
			return
		}
	}
}
