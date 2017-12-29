package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	go func(message string) {
		fmt.Println("GO GO GO !!!!")
		//<-messages

	}("ping!")
	//messages <- "message"
	<-messages
	fmt.Println("!!!!!")
}
