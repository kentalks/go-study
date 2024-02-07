package main

import (
	"math/rand"
	"time"

	"kensoft.tech/go-study/function"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

}

func main() {
	function.HelloSlice()

	//function.HelloArray()

	/*
		cat := function.Cat{
			Age: 0,
			Pet: &function.Pet{Name: "小白"},
		}
		function.Super(&cat)
		(&cat).Super()
		cat.Super()
		cat.Info()
		fmt.Println(cat.Pet)
	*/

	//function.PrintHellos([]string{"mei", "ken"})

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

	/*
		fmt.Println(function.GCD(-6, 18))
	*/

}
