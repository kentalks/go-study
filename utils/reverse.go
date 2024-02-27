package function

import "strconv"

func ReverseInt(i int) int {
	i, _ = strconv.Atoi(string(strconv.Itoa(i)))
	return i
}
