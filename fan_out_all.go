package main

func FanOutAll(input <-chan string, outs... chan<- string) {
	go func() {
		for {
			msg := <-input
			for _, out := range outs {
				go func(o chan<-string) {
					o <- msg
				}(out)
			}
		}
	}()
}