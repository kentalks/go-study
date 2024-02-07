package function

import "fmt"

type Pet struct {
	Name string
}

func (p *Pet) Info() {
	fmt.Printf("I'm %s.\n", p.Name)
}

type Cat struct {
	*Pet
	Age    int
	Father *Cat
}

func NewCat(name string, age int) Cat {
	return Cat{
		Pet: &Pet{name},
		Age: 3,
	}
}

func Super(s *Cat) {
	s.Age += 1

}

func (s *Cat) Super() {
	s.Age += 2
}

type Point struct {
	X int
	Y int
}
