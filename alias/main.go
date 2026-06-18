package main

import "fmt"

func main() {
	var i mile
	fmt.Println(i)
	i += 12
	fmt.Println(i)

	var intOperation IntOperation = add
	applyOperation(12, 12, intOperation)

}

type mile int

type IntOperation func(int, int) int

func applyOperation(a int, b int, op IntOperation) {
	result := op(a, b)
	fmt.Println(result)
}
func add(a, b int) int {
	return a + b
}
