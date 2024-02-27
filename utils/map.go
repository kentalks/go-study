package function

import "fmt"

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func HelloMap() {
	lookup := make(map[string]int)
	lookup["goku"] = 100
	p1, p2 := lookup["none"]
	fmt.Println(p1, p2)

	lookup = map[string]int{
		"goku":  100,
		"gohan": 200,
	}

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{
		Name: "krillin",
	}
	fmt.Println(goku.Friends["krillin"])
}
