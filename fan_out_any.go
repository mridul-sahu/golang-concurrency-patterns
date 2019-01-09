package main

func FanOutAny(input <-chan string, out1, out2 chan<- string) {
	go func() {
		for {
			msg := <-input
			select {
			case out1 <- msg:
			case out2 <- msg:
			}
		}
	}()
}
