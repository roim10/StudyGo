package main

import "fmt"

func main() {
	var i *int
	fmt.Println(i)
	var k int = 12
	i = &k
	fmt.Println(i)
	fmt.Println(*i)

	p := new(int)
	*p = k
	p = &k
	fmt.Println(*p)

	var p_1 **int
	p_1 = &p
	fmt.Println(p_1)

	var array [5]*int
	var a, b, n, d int = 12, 12, 12, 12
	array[1] = &b
	array[2] = &a
	array[4] = &d
	array[3] = &n
	fmt.Println(array)

	for _, v := range array {
		if v != nil {
			fmt.Println(v)
		} else {
			fmt.Println("lol")
		}
	}
}
