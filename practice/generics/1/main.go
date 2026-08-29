package main

import (
	"fmt"
	"strconv"
)

func Filter[T any](items []T, predicate func(T) bool) []T {
	resulte := make([]T, 0)
	for _, v := range items {
		if predicate(v) {
			resulte = append(resulte, v)
		}
	}
	return resulte
}
func Map[T, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = transform(v)
	}
	return result
}
func Reduce[T, U any](items []T, initial U, accumulate func(U, T) U) U {
	result := initial
	for _, v := range items {
		result = accumulate(result, v)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	even := Filter(nums, func(n int) bool { return n%2 == 0 })
	// [2 4 6]

	doubled := Map(nums, func(n int) int { return n * 2 })
	// [2 4 6 8 10 12]

	strs := Map(nums, func(n int) string { return strconv.Itoa(n) })
	// ["1" "2" "3" "4" "5" "6"] — здесь T=int, U=string

	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	// 21

	fmt.Println(even)
	fmt.Println(doubled)
	fmt.Println(strs)
	fmt.Println(sum)
}
