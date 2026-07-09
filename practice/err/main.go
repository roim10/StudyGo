package main

import "fmt"

var prices = map[int]float64{
	1: 299.99,
	2: 49.90,
	3: 1200.00,
}

func getPrice(productID int) (float64, any) {
	price, exists := prices[productID]
	if !exists {
		return 0, "Товар с таким ID не найден"
	}
	return price, nil
}

func calcTotal(price float64, quantity int) (float64, any) {
	if quantity <= 0 || quantity > 100 {
		return 0, "больше нуля покупок должно быть и не превышать 100"
	}
	if price <= 0 {
		return 0, "цена не может быть равна 0 или быть меньше 0"
	}
	var result float64 = price * float64(quantity)
	return result, nil
}

func applyDiscount(total float64, discountPercent int) (float64, any) {
	if discountPercent > 50 || discountPercent < 0 {
		return 0, "неправильно"
	}
	var percent float64 = float64(discountPercent) / 100
	var result float64 = total - (total * percent)
	return result, nil
}

func main() {
	fact, err := getPrice(3)
	fmt.Println(fact)
	if err != nil {
		fmt.Printf("ошибка %v\n", err)
		return
	}
	fact, err = calcTotal(299.99, 3)
	fmt.Println(fact)
	if err != nil {
		fmt.Printf("ошибка %v\n", err)
		return
	}
	fact, err = applyDiscount(fact, 20)
	fmt.Println(fact)
	if err != nil {
		fmt.Printf("ошибка %v\n", err)
		return
	}
}
