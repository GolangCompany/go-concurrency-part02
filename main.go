package main

import (
	"fmt"
)

func main() {

	msgch := make(chan string, 128)

	msgch <- "A"
	msgch <- "B"
	msgch <- "C"

	Rmsg := <-msgch

	fmt.Println(Rmsg)

}
