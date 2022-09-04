package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// function.PrintHellos([]string{"mei", "ken"})

	// r := function.RandString(10)
	// fmt.Println(r)
	// fmt.Println(function.ToUpper(r))

	// s := function.RandSeq(10)
	// fmt.Println(s)

}
