package main

func BoringFanIn(chans ... <-chan string) <-chan string {
	c := make(chan string)
	for _, v := range chans {
		go func(<-chan string) {
			for {
				c <- <-v
			}
		}(v)
	}
	return c
}
