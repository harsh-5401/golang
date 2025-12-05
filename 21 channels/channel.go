package main

import "fmt"

func main() {
	messageChan := make(chan string)

	messageChan <- "ping" // sending inside channel

	msg := <-messageChan // getting data fromc channel

	fmt.Println(msg)

}