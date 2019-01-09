package main

import "fmt"

type Message struct {
	msg  string
	wait chan bool
}

func BoringSyncedGenerator(msg string) <-chan Message {
	m := make(chan Message)

	go func() {
		for i := 0; ; i++ {
			wait := make(chan bool)
			m <- Message{msg: fmt.Sprintf("%s %d", msg, i), wait: wait}
			<-wait
		}
	}()
	return m
}

func BoringSequentialFanIn(syncedChans ... <-chan Message) <-chan string {
	c := make(chan string)

	go func() {
		var msgs []Message

		for _, v := range syncedChans {
			msgs = append(msgs, <-v)
		}

		for _, v := range msgs {
			c <- v.msg
		}

		for _, v := range msgs {
			v.wait <- true
		}
	}()

	return c
}
