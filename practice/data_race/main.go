package main

import (
	"fmt"
	"sync"
	"time"
)

// VideoStat хранит аналитику по конкретному видео
type VideoStat struct {
	VideoID string
	Views   int
}

func main() {
	// Создаем объект статистики для видео
	stat := &VideoStat{
		VideoID: "advanced-go-concurrency",
		Views:   0,
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	usersCount := 1000
	// Симулируем 1000 одновременных запросов от пользователей
	for i := 0; i < usersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 1)
			mu.Lock()
			stat.Views++
			mu.Unlock()
		}()

	}
	// Ждем, пока все горутины завершат работу
	wg.Wait()
	fmt.Printf("Ожидаемое количество просмотров: %d\n", usersCount)
	fmt.Printf("Фактическое количество просмотров: %d\n", stat.Views)
}
