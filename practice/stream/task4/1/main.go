package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("numbers.txt", os.O_RDWR|os.O_CREATE, 0644)
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for i := range 21 {
		writer.WriteString(strconv.Itoa(i))
		writer.WriteString(" ")
	}
	writer.Flush()
}
