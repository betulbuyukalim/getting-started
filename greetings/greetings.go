package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting() string {
	greeting := []string{
		"hi %v! good to see you again!",
		"hey %v! you look good!",
		"you're doing well. proud of you, %v!",
	}

	return greeting[rand.Intn(len(greeting))]
}

func Hellos(names []string)(map[string]string, error){
	messages := make(map[string]string)
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
