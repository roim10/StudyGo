package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("words.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	file1, err1 := os.OpenFile("long_words.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	defer file1.Close()
	var str string
	for {
		line, err := reader.ReadString('\n')
		str += line
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer := bufio.NewWriter(file1)
	words := strings.Fields(str)
	for _, v := range words {
		if len(v) > 5 {
			writer.WriteString(v + "\n")
		}
	}
	writer.Flush()
}
