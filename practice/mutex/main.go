package main

import "sync"

type ProductCache struct {
	values map[string]float64
	mu     sync.RWMutex
}

func (pr *ProductCache) Set(id string, price float64, wg *sync.WaitGroup) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	defer wg.Done()
	pr.values[id] = price
}

func (pr *ProductCache) Get(id string, wg *sync.WaitGroup) (float64, bool) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	defer wg.Done()
	price, ok := pr.values[id]
	return price, ok
}

func (pr *ProductCache) Delete(id string, wg *sync.WaitGroup) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	defer wg.Done()
	delete(pr.values, id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3 * 5)
	var prodectCache ProductCache = ProductCache{values: make(map[string]float64)}
	for i := 1; i <= 5; i++ {
		go prodectCache.Set("lol", float64(i), &wg)
		go prodectCache.Get("lol", &wg)
		go prodectCache.Delete("lol", &wg)
	}
	wg.Wait()
}
