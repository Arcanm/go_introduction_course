package main

import "fmt"

// Function that only SENDS data to the channel
func say(text string, c chan<- string) {
	c <- text // Sends "Bye" to the channel
}

// Function that only RECEIVES data from the channel
func receive(c <-chan string) {
	fmt.Println(<-c) // Receives from the channel and prints it
}

func main() {
	// Create a buffered channel with capacity 1
	channel := make(chan string, 1)

	fmt.Println("Hello")

	// Start the goroutine to send "Bye" to the channel
	go say("Bye", channel)

	// Start another goroutine to receive and print the value
	go receive(channel)

	// Wait for user input to allow goroutines to finish
	fmt.Scanln()
}
