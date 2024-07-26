package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int64

func mainWrong() {
	wg := new(sync.WaitGroup)
	for x := 1; x <= 30; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := 1; y <= 100_000; y++ {
				i++
			}
		}()
	}
	wg.Wait()
}

func mainRightAtomic() {
	wg := new(sync.WaitGroup)
	for x := 1; x <= 30; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := 1; y <= 100_000; y++ {
				atomic.AddInt64(&i, 1)
			}
		}()
	}
	wg.Wait()
}

func mainRightLock() {
	wg := new(sync.WaitGroup)
	lock := new(sync.Mutex)
	for x := 1; x <= 30; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := 1; y <= 100_000; y++ {
				lock.Lock()
				i++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
}

func main() {
	mainRightLock()
	fmt.Println(i)
}
