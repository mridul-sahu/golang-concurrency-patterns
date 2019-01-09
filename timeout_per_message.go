package main

import (
	"fmt"
	"time"
)

func BoringTimeoutPerMessage(input <-chan string) {
	for {
		select {
		case s := <-input:
			fmt.Print(s)
		case <-time.After(time.Second * 3):
			fmt.Print("Too boring, I am leaving")
			return
		}
	}
}
