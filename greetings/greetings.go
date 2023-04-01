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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
        
		// Associate the retrieved message with the name.
		messages[name] = message

	}
	return messages, nil
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
