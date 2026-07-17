package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrInvalidPrice = errors.New("цена должна быть положительной")
var ErrInvalidDiscount = errors.New("скидка должна быть от 0 до 100")

type negative_number_error struct{}

func (error_obj negative_number_error) Error() string {
	return "Invalid parameter"
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		obj := negative_number_error{}
		return 0, obj
	}
	sqrt := math.Sqrt(n)
	return sqrt, nil
}
func validateAge(age int) (bool, error) {
	switch {
	case age < 0:
		return false, errors.New("возраст не может быть отрицательным")
	case age > 125:
		return false, errors.New("возраст нереалистично большой")
	default:
		return true, nil
	}
}

func calculateDiscount(price float64, discountPercent int) (float64, error) {
	NewPrice := price * (1 - float64(discountPercent)/100)
	//switch {
	//case price <= 0:
	//	return 0, errors.New("цена ниже или равна 0")
	//case discountPercent < 0 || discountPercent > 100:
	//	return 0, errors.New("проценты не могут быть ниже 0 и больше 100")
	//default:
	//	return NewPrice, nil
	//}
	switch {
	case price <= 0:
		return 0, ErrInvalidPrice
	case discountPercent < 0 || discountPercent > 100:
		return 0, ErrInvalidDiscount
	default:
		return NewPrice, nil
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	result := a / b
	return result, nil
}

func main() {
	fact, err := divide(5, 5)
	fmt.Println(fact)
	fmt.Println(err)
	fact, err = divide(5, 0)
	fmt.Println(fact)
	fmt.Println(err)
	_, err1 := calculateDiscount(-10, 20)
	if errors.Is(err1, ErrInvalidPrice) {
		fmt.Println("Это ошибка именно про цену")
	}
	validateAge(127)
	validateAge(-2)
	validateAge(121)
	sqrt(12)
	sqrt(-4)
}
