package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings err: ")
	log.SetFlags(0)

	// get greeting msg and print it
	msg, err := greetings.Hello("Tori")
	// exit program if error is returned
	if err != nil {
		log.Fatal(err)
	}

	// Print message if no errors returned
	fmt.Println(msg)
}
