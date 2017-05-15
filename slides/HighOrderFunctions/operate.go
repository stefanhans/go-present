package main

import "fmt"

// OPERATOR START OMIT
type BinaryOperator func(x, y int) int

func (op BinaryOperator) Operate(x, y int) int { return op(x, y) }
// OPERATOR END OMIT

func main() {

	var mult BinaryOperator = func(i, j int) int { return i*j}

	fmt.Println(mult.Operate(3, 5))
}