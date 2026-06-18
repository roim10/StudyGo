package main

import "fmt"

func main() {
	var people map[string]int //Карта определяется как объект типа map[K]V, где К представляет тип ключа, а V - тип значения
	fmt.Println(people)
	people1 := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people1)
	people1["Bob"] = 6
	fmt.Println(people1)
	if val, ok := people1["Tom"]; ok {
		fmt.Println(val)
	}
	for i, v := range people1 {
		fmt.Println(i, v)
	}
	people2 := make(map[string]int)
	fmt.Println(people2)
	people2["kate"] = 12
	delete(people2, "Sam") //Для удаления применяется встроенная функция delete(map, key), первым параметром
	fmt.Println(people2)   //которой является карта, а вторым - ключ, по которому надо удалить элемент.

}
