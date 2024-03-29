package utils

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func SortName() {
	people := []Person{
		{"xzan", "kan", 20},
		{"yyy", "zzz", 30},
		{"xxx", "aaa", 40},
	}

	sort.Slice(people, func(i int, j int) bool {
		return people[i].FirstName > people[j].FirstName
	})
	fmt.Println(people)
}

func MakeMult(base int) func(int) int {
	return func(i int) int { return base * i }
}
