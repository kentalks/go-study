package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func HelloSlice() {
	scores := []int{1, 2, 3}
	scores2 := make([]int, 10)
	scores3 := make([]int, 0, 5)
	scores3 = scores3[0:3]
	scores3 = append(scores3, 1)

	fmt.Println(len(scores), cap(scores))
	fmt.Println(len(scores2), cap(scores2))
	fmt.Println(len(scores3), cap(scores3))

	haystack := "the spice must flow"
	fmt.Println(strings.Index(haystack[5:], " ")) //找到第5个字符之后的第一个空格

	fmt.Println(haystack[0 : len(haystack)-1])

	scores = removeAtIndex(scores, 0)
	fmt.Println(scores)

	copyTo()

}

// 从乱序的切片中去除一个值
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func copyTo() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores) //排序
	worst := make([]int, 5)
	copy(worst, scores[:5])

	fmt.Println(worst)
}
