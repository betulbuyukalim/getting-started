package main

import (
	"fmt"
	"log"

	"github.com/betulbuyukalim/go-site/getting-started/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"betul", "sevim", "mithat"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
	//fmt.Println(messages)

	mes, err := greetings.Hello("betul")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mes)
}
