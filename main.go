package main

import (
	"fmt"

	"kensoft.tech/go-study/function"
)

func main() {
	mult2 := function.MakeMult(2)
	mult5 := function.MakeMult(5)
	fmt.Println(mult2(10), mult5(10))

}
