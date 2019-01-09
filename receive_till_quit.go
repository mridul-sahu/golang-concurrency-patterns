package main

import "fmt"

func BoringReceiveTillQuit(input <-chan string, quit <-chan bool) {
	for {
		select {
		case s := <-input:
			fmt.Print(s)
		case <-quit:
			fmt.Print("Oh ok that was nice talk")
			return
		}
	}
}
