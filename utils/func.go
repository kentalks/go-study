package function

import (
	"errors"
	"fmt"
)

func Div(num, den int) (res int, rem int, err error) {
	if den == 0 {
		return 0, 0, errors.New("devide 0")
	}
	return num / den, num % den, nil
}

func AddTo(base int, vals ...int) []int {
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
