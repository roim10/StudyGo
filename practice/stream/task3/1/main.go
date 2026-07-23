package main

import (
	"fmt"
	"io"
	"os"
)

type Student struct {
	name         string
	age          int
	averageScore float64
}

func main() {
	Tom := Student{
		name:         "Tom",
		age:          18,
		averageScore: 4.87,
	}
	ALex := Student{
		name:         "Alex",
		age:          17,
		averageScore: 4.30,
	}
	Lin := Student{
		name:         "Lin",
		age:          19,
		averageScore: 3.08,
	}
	file, err := os.OpenFile("students.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "%-10s %-10d %-10.2f\n", Tom.name, Tom.age, Tom.averageScore)
	fmt.Fprintf(file, "%-10s %-10d %-10.2f\n", ALex.name, ALex.age, ALex.averageScore)
	fmt.Fprintf(file, "%-10s %-10d %-10.2f\n", Lin.name, Lin.age, Lin.averageScore)
	var (
		name         string
		age          int
		averageScore float64
	)
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		_, err := fmt.Fscanf(file, "%s %d %f\n", &name, &age, &averageScore)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, averageScore)
	}

}
