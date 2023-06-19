package main

import (
	"fmt"
)

func main() {

	msgch := make(chan string, 128)

	msgch <- "A"
	msgch <- "B"
	msgch <- "C"

	for Rmsg := range msgch {
		fmt.Println(Rmsg)
	}

}

//to run perfect uncommented the code and comment the above code

// package main
// import (
//     "fmt"
// )

// func main() {
//     msgch := make(chan string, 128)
//     msgch <- "A"
//     msgch <- "B"
//     msgch <- "C"
//     close(msgch)
//     for Rmsg := range msgch {
//         fmt.Println(Rmsg)
//     }
// }
