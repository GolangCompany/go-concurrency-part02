package main

import (
	"fmt"
)

func main() {

	msgch := make(chan string, 128)

	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)
	for {
		Rmsg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println(Rmsg)
	}
}
