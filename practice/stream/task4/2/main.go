package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var sum int
	for {
		line, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		line = strings.TrimSpace(line)
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += n
	}
	fmt.Println(sum)
}
