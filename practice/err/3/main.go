package main

import "fmt"

func SafeCalculate(a, b int, op string) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("восстановились после паники: %v", r)
		}
	}()

	switch op {
	case "/":
		result = a / b
		return
	case "*":
		result = a * b
		return
	case "+":
		result = a + b
		return
	case "-":
		result = a - b
		return
	case "%":
		result = a % b
		return
	default:
		panic("неизвестная операция: " + op)
	}
}

func main() {
	result, err := SafeCalculate(10, 2, "/")
	fmt.Println(result)
	fmt.Println(err)
	// result = 5, err = nil

	result, err = SafeCalculate(10, 0, "/")
	fmt.Println(result)
	fmt.Println(err)
	// result = 0, err = "восстановились после паники: runtime error: integer divide by zero"

	result, err = SafeCalculate(5, 3, "^")
	fmt.Println(result)
	fmt.Println(err)
	// result = 0, err = "восстановились после паники: неизвестная операция: ^"
}
