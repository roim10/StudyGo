package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func downloadFile(name string, size int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Началась загрузка файла %v (%v MB)...\n", name, size)
	time.Sleep(time.Duration(size) * time.Microsecond)
	fmt.Printf("Файл %s успешно загружен!\n", name)
}

func main() {
	var format []string = []string{"image.png", "video.mp4", "document.pdf", "database.db"}
	var size []int = []int{100, 500, 50, 1200}
	wg.Add(4)
	//for i := range format {
	//	go downloadFile(format[i], size[i], &wg)
	//}
	//wg.Wait()
	//fmt.Println("Все файлы успешно скачаны. Менеджер завершил работу.")
	for i := range format {
		go func(name string, size int) {
			defer wg.Done()
			fmt.Printf("Началась загрузка файла %v (%v MB)...\n", name, size)
			time.Sleep(time.Duration(size) * time.Microsecond)
			fmt.Printf("Файл %s успешно загружен!\n", name)
		}(format[i], size[i])
	}
	wg.Wait()
}
