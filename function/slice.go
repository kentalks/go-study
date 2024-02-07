package function

import "fmt"

func HelloSlice() {
	scores := []int{1, 2, 3}
	scores2 := make([]int, 10)
	scores3 := make([]int, 0, 5)
	scores3 = scores3[0:3]
	scores3 = append(scores3, 1)

	fmt.Println(len(scores), cap(scores))
	fmt.Println(len(scores2), cap(scores2))
	fmt.Println(len(scores3), cap(scores3))
}
