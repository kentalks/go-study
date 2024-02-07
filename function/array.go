package function

import "fmt"

func HelloArray() {
	scores := [5]int{10, 10}
	scores[1] = 20

	for i, v := range scores {
		if i == 2 && v == 0 {
			scores[i] = 30
		}
	}
	fmt.Println(scores)
}
