package main

import (
	"fmt"
	"time"
)

// Function that sends a message into a channel
func message(text string, c chan string) {
	c <- text // Send the text to the channel
}

func main() {
	// Create a buffered channel with a capacity of 2
	c := make(chan string, 2)

	// Send messages to the channel
	c <- "FirstMessage"
	c <- "SecondMessage"

	// Print the current length and capacity of the channel
	fmt.Println("Channel length:", len(c), "Capacity:", cap(c))

	// --- Closing a Channel ---
	// Closing a channel means no more values can be sent, but you can still receive remaining values.
	close(c)

	// This would cause a runtime error because the channel is closed
	// c <- "AnotherMessage"

	// --- Iterating Over a Channel ---
	// Using 'range' to receive values until the channel is empty
	for message := range c {
		fmt.Println("Received from channel:", message)
	}

	// --- Using Select with Multiple Channels ---
	message1 := make(chan string)
	message2 := make(chan string)

	// Launch two goroutines to send messages asynchronously
	go message("Hello", message1)
	go message("Hello 2", message2)

	// Using 'select' to receive from multiple channels
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-message1:
			fmt.Println("Message 1 received:", m1)
		case m2 := <-message2:
			fmt.Println("Message 2 received:", m2)
		}
	}

	// --- Advanced Example: Synchronizing Goroutines Using Channels ---
	advancedChannelExample()
}

// Advanced example: Using channels to coordinate multiple goroutines
func advancedChannelExample() {
	// Create a channel for integer values
	numbers := make(chan int, 3)

	// Goroutine to produce values
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Sending:", i)
			numbers <- i // Send value to channel
			time.Sleep(time.Millisecond * 500)
		}
		close(numbers) // Close channel when done
	}()

	// Receiving values from the channel
	for num := range numbers {
		fmt.Println("Received:", num)
	}
}
