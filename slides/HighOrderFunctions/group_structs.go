package main

import (
	"fmt"
)

// DEF_START OMIT
type Person struct {
	Name string
	Age  int
	Male bool
}

type ListOfPerson []Person

type PersonsGroupByMap map[interface{}]int

type PersonsGroupByFunc func(p Person) (interface{}, int)

func (list ListOfPerson) GroupBy(f PersonsGroupByFunc) PersonsGroupByMap {
	var out PersonsGroupByMap = make(map[interface{}]int)

	for _, p := range list {
		i, n := f(p)
		out[i] = out[i] + n
	}
	return out
}

// DEF_END OMIT

func main() {

	// STRUCT_START OMIT
	list := ListOfPerson{
		Person{Name: "Peter", Age: 18, Male: true},
		Person{Name: "Petra", Age: 21, Male: false},
		Person{Name: "Karl", Age: 21, Male: true},
		Person{Name: "Gustav", Age: 21, Male: true},
		Person{Name: "Sabine", Age: 22, Male: false},
		Person{Name: "Sabine", Age: 22, Male: false},
		Person{Name: "Stefan", Age: 24, Male: true},
		Person{Name: "Jochen", Age: 25, Male: true},
		Person{Name: "Eva", Age: 21, Male: false},
		Person{Name: "Peter", Age: 18, Male: true},
		Person{Name: "Petra", Age: 21, Male: false},
		Person{Name: "Karl", Age: 21, Male: true},
		Person{Name: "Gustav", Age: 21, Male: true},
		Person{Name: "Sabine", Age: 22, Male: false},
		Person{Name: "Sabine", Age: 22, Male: false},
		Person{Name: "Stefan", Age: 24, Male: true},
		Person{Name: "Jochen", Age: 25, Male: true},
		Person{Name: "Eva", Age: 21, Male: false},
		Person{Name: "Peter", Age: 18, Male: true},
		Person{Name: "Petra", Age: 21, Male: false},
		Person{Name: "Karl", Age: 21, Male: true},
		Person{Name: "Gustav", Age: 21, Male: true},
	}
	// STRUCT_END OMIT

	// RUN_START OMIT
	countByMale := func(p Person) (interface{}, int) { return p.Male, 1 }
	fmt.Printf("list.GroupBy(countByMale): %v\n", list.GroupBy(countByMale))

	countByAge := func(p Person) (interface{}, int) { return p.Age, 1 }
	fmt.Printf("list.GroupBy(countByAge): %v\n", list.GroupBy(countByAge))

	countByName := func(p Person) (interface{}, int) { return p.Name, 1 }
	fmt.Printf("list.GroupBy(countByName): %v\n", list.GroupBy(countByName))
	// RUN_END OMIT
}

/*


	sumAgeByMale := func(p Person) (interface{}, int) { return p.Male, p.Age }
	fmt.Printf("list.GroupBy(sumAgeByMale): %v\n", list.GroupBy(sumAgeByMale))
*/
