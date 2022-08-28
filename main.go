package main

import (
	"fmt"
	"log"

	"kensoft.tech/go-lib/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"ken", "mei", "jia", "tai"}
	ms, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ms {
		fmt.Println(v)
	}
}
