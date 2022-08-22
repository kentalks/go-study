package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	opMain()
	fmt.Println()
}

func div2(num, den int) (res int, rem int, err error) {
	if 0 == den {
		return 0, 0, errors.New("devide 0")
	}
	return num / den, num % den, nil
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

type FuncOpts struct {
	First string
	Last  string
	Age   int
}

func MyFunc(opts FuncOpts) error {
	fmt.Println(opts.First, opts.Last, opts.Age)
	return nil
}

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

// Define a function type
type opFunc func(int, int) int

var opMap = map[string]opFunc{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func opMain() {
	exps := [][]string{
		{"2", "+", "2"},
		{"2", "-", "2"},
		{"2", "/", "2"},
		{"2", "*", "2"},
		{"2", "%", "2"},
		{"a", "+", "b"},
		{"a"},
	}

	for _, e := range exps {
		if len(e) != 3 {
			fmt.Println("invald:", e)
			continue
		}
		p1, err := strconv.Atoi(e[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := e[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported：", op)
			continue
		}
		p2, err := strconv.Atoi(e[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
