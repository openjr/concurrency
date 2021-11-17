package concurrency

import (
	"sync"
)

func AddWithChannels(concurrency int) int {
	var sum int
	sumChannel := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			select {
			case i :=  <- sumChannel:
				sum += i
			case <- done:
				return
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(concurrency)

	for i:=0; i< concurrency; i++ {
		go func(i int, wg *sync.WaitGroup) {
			sumChannel <- i
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
	done <- true
	return sum
}

func AddWithMutex(concurrency int) int {
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	m := sync.RWMutex{}
	var sum int
	for i := 0; i < concurrency; i++ {
		go func(group *sync.WaitGroup, i int) {
			m.Lock()
			sum += i
			m.Unlock()
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	return sum
}