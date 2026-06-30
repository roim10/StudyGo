package main

import (
	"fmt"
	"sync"
	"time"
)

type MetricStorage struct {
	MetricStorage map[string]int
	Rmu           sync.RWMutex
}

func (ms *MetricStorage) PrintResults() {
	ms.Rmu.RLock() // Используем RLock для чтения
	defer ms.Rmu.RUnlock()

	fmt.Println("\n--- Итоговая статистика ---")
	for service, count := range ms.MetricStorage {
		fmt.Printf("Сервис [%s] обработал запросов: %d\n", service, count)
	}
}

func (ms *MetricStorage) fetchMetrics(serviceName string, metricsChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		time.Sleep(300 * time.Millisecond)

		// Имитируем сбор метрики (например, получили i-количество запросов)
		metricValue := i + 1

		// Блокируем ТОЛЬКО операцию записи в общую map
		ms.Rmu.Lock()
		ms.MetricStorage[serviceName] += metricValue
		ms.Rmu.Unlock()

		// Отправляем в канал тоже вне мьютекса, чтобы не блокировать другие воркеры
		metricsChan <- metricValue
	}
}

func main() {
	var wg sync.WaitGroup
	metricsChan := make(chan int)
	wg.Add(3)
	MetricStorage := &MetricStorage{
		MetricStorage: make(map[string]int),
	}
	go MetricStorage.fetchMetrics("lol", metricsChan, &wg)
	go MetricStorage.fetchMetrics("pop", metricsChan, &wg)
	go MetricStorage.fetchMetrics("kek", metricsChan, &wg)
	go func() {
		wg.Wait()
		close(metricsChan)
	}()
	//for metric := range metricsChan {
	//fmt.Println("Получено:", metric)
	//}
mainLoop:
	for {
		select {
		case metric, ok := <-metricsChan:
			if !ok {
				break mainLoop
			}

			fmt.Println(metric)

		case <-time.After(3 * time.Second):
			fmt.Println("Warning: Data timeout, waiting for services...")
		}
	}
	MetricStorage.PrintResults()
}
