package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var i = []string{"k", "kd", "fd"}
	j := make([]string, 12)
	u := i[1:2]
	fmt.Println(j, i, u)
	usr := []int{1, 2, 2, 3}
	usr1 := usr[1]
	fmt.Println(usr1)
	for _, v := range usr {
		fmt.Println(v)
	}
	fmt.Println(cap(usr))
	d2 := [][]int{
		[]int{1, 3},
		[]int{2, 4},
		[]int{4, 7},
	}
	fmt.Println(d2)
	usr = append(usr, 12, 12)
	delet := []int{2, 3, 4, 45, 6, 4}
	var n int = 3
	delet = append(delet[:n], delet[n+1:]...)
	fmt.Println(delet)
	cop := make([]int, 8)
	copy(cop, usr)
	fmt.Println(cop)
	sort.Ints(usr)
	fmt.Println(usr)
	sort.SearchInts(usr, 12)
	fmt.Println(usr)
	fmt.Println(reflect.DeepEqual(usr, delet))
}
