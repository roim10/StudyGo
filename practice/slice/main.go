package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	inventory := []string{"меч", "зелье", "щит"}
	inventory = addItem(inventory, "шлем")
	inventory = removeItem(inventory, 2)
	fmt.Println(inventory)
	scores := []int{45, 12, 89, 33, 27, 50, 91, 102, 5}
	sortedScores, targetIndex, topScores := Sort(scores)
	fmt.Println("Отсортированные (reverse) очки:", sortedScores)
	fmt.Println("Индекс числа 33:", targetIndex)
	fmt.Println("Топ-3 игрока:", topScores)
	originalBackup := []string{"System", "Data", "Config"}
	originalBackup, lol := srav(originalBackup)
	fmt.Println(originalBackup, lol)
	board := [][]string{
		[]string{"0", "1", "2"},
		[]string{"0", "1", "2"},
		[]string{"0", "1", "2"},
	}
	board = matrix(board)
}

func addItem(s []string, item string) []string {
	s = append(s, item)
	return s
}

func removeItem(v []string, index int) []string {
	v = append(v[:index], v[index+1:]...)
	return v
}
func Sort(scores1 []int) ([]int, int, []int) {
	// 1. Сортируем по возрастанию (нужно для корректного поиска)
	sort.Ints(scores1)
	// 2. Ищем индекс числа 33
	index := sort.SearchInts(scores1, 33)
	// 3. Переворачиваем (теперь самые большие числа в начале)
	sort.Sort(sort.Reverse(sort.IntSlice(scores1)))
	// 4. Берем топ-3 (первые три элемента после переворота)
	top3 := scores1[:3]
	return scores1, index, top3
}
func srav(originalBackup1 []string) ([]string, bool) {
	curentSession := make([]string, len(originalBackup1))
	copy(curentSession, originalBackup1)
	curentSession[1] = "all"
	l := reflect.DeepEqual(curentSession, originalBackup1)
	return originalBackup1, l
}
func matrix(board1 [][]string) [][]string {
	for _, m := range board1 {
		fmt.Println(m)
	}
	return board1
}
