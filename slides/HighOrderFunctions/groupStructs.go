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

type MapGroupByPerson map[interface{}]int

type PersonGroupFunc func(p Person) (interface{}, int)

func (list ListOfPerson) GroupBy(f PersonGroupFunc) MapGroupByPerson {
	var out MapGroupByPerson = make(map[interface{}]int)

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
		Person{ Name: "Peter",   Age:  18, Male: true,  },
		Person{ Name: "Petra",   Age:  21, Male: false, },
		Person{ Name: "Karl",    Age:  21, Male: true,  },
		Person{ Name: "Gustav",  Age:  21, Male: true,  },
		Person{ Name: "Sabine",  Age:  22, Male: false, },
		Person{ Name: "Sabine",  Age:  22, Male: false, },
		Person{ Name: "Stefan",  Age:  24, Male: true,  },
		Person{ Name: "Jochen",  Age:  25, Male: true,  },
		Person{ Name: "Eva",     Age:  21, Male: false, },
	}
	// STRUCT_END OMIT

	//fmt.Printf("list: %v\n", list)

	// RUN_START OMIT
	countByMale := func(p Person) (interface{}, int) {
		return p.Male, 1
	}
	countByMaleMap := list.GroupBy(countByMale)
	fmt.Printf("list.GroupBy(countByMale): %v\n", countByMaleMap)

	sumAgeByMale := func(p Person) (interface{}, int) {
		return p.Male, p.Age
	}
	sumAgeByMaleMap := list.GroupBy(sumAgeByMale)
	fmt.Printf("list.GroupBy(sumAgeByMale): %v\n", sumAgeByMaleMap)

	countByAge := func(p Person) (interface{}, int) {
		return p, 1
	}
	countByAgeMap := list.GroupBy(countByAge)
	fmt.Printf("list.GroupBy(countByAge): %v\n", countByAgeMap)

	countByName := func(p Person) (interface{}, int) {
		return p.Name, 1
	}
	countByNameMap := list.GroupBy(countByName)
	fmt.Printf("list.GroupBy(countByName): %v\n", countByNameMap)

	// RUN_END OMIT
}
