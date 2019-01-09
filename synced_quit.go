package main

import "fmt"

func BoringSyncedQuit (input <-chan string, quit chan bool) {
	for {
		select {
		case s := <-input: fmt.Print(s)
		case <-quit:
			fmt.Print("Well see ya, bye")
			quit <- true
		}
	}
}
