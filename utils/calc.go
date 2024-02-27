package utils

import (
	"fmt"
	"strconv"
)

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

func Calc() {
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

// 欧几里得最大公约数算法
func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	c := a % b
	return GCD(b, c)

}
