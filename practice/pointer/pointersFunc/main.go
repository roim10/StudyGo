package main

import "fmt"

func main() {
	in := 12
	Fsquar(in)
	fmt.Println(in)
	squar(&in)
	fmt.Println(in)
}

func squar(ix *int) *int {
	*ix = *ix * *ix
	return ix
}
func Fsquar(iv int) int {
	iv = iv * iv
	return iv
}
