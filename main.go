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

	// fmt.Println(function.ProcessStr("asdfasdf", 12))

	/*
		in := [3]string{"a", "b", "c"}
		var out []*string
		for _, v := range in {
			v := v
			out = append(out, &v)
		}
		fmt.Println(*out[0], *out[1], *out[2])
	*/

}
