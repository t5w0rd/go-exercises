package exercises

import (
	"sync"
	"sync/atomic"
)

func ConcurrencyAdd(addr *int64, threads int, loops int) {
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < loops; i++ {
				*addr++
			}
		}()
	}
	wg.Wait()
}

func ConcurrencyAdd2(addr *int64, threads int, loops int) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < loops; i++ {
				mtx.Lock()
				*addr++
				mtx.Unlock()
			}
		}()
	}
	wg.Wait()
}

func ConcurrencyAdd3(addr *int64, threads int, loops int) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1000)
	go func() {
		defer wg.Done()
		for range ch {
			*addr++
		}
	}()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < loops; i++ {
				ch <- struct{}{}
			}
		}()
	}
	wg.Wait()

	wg.Add(1)
	close(ch)
	wg.Wait()
}

func ConcurrencyAdd4(addr *int64, threads int, loops int) {
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < loops; i++ {
				atomic.AddInt64(addr, 1)
			}
		}()
	}
	wg.Wait()
}
