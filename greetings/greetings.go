package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Returns greeting for named person
func Hello(name string) (string, error) {
	// Handle error if no name is given
	if name == "" {
		return "", errors.New("No name provided in greeting function.")
	}
	// return greeting which embeds name in message
	message := fmt.Sprintf(randomFormat(), name)
	// Return message and an error value of nil
	return message, nil
}

// init sets initial values for variables used in the function.
// go executes the init funciton automatically at startup (once global variables have been initialized)
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
