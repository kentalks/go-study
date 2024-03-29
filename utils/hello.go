package utils

import (
	"fmt"
	"log"

	util "kensoft.tech/go-lib"
)

func PrintHellos(names []string) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	ms, err := util.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ms {
		fmt.Println(v)
	}
}
