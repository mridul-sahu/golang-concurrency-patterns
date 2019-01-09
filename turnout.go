package main

func Turnout(inp1, inp2 <-chan string, out1, out2 chan<- string) {
	c := BoringFanInFixedArguments(inp1, inp2)
	FanOutAny(c, out1, out2)
}
