package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
	 // If no name was given, return an error with a message.
	 if name == "" {
        return "", errors.New("empty name")
    }

	message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string);
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	//fmt.Print("init greet function")
	//fmt.Print(time.Now().UnixNano())
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome",
		"Hey %v, How many words?",
		"Its ok %v, dekho kya hota hai?",
	}

	return formats[rand.Intn(len(formats))]
}