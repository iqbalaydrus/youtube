package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Summary struct {
	ValSum   float64
	ValCount int32
	ValMin   float64
	ValMax   float64
}

func main() {
	goConcurrentAlg()
}

func basic() {
	now := time.Now()
	f, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	result := map[string]Summary{}
	for scanner.Scan() {
		sSplit := strings.Split(scanner.Text(), ";")
		val, err := strconv.ParseFloat(sSplit[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		summmary := result[sSplit[0]]
		summmary.ValSum += val
		summmary.ValCount++
		if val > summmary.ValMax {
			summmary.ValMax = val
		}
		if val < summmary.ValMin {
			summmary.ValMin = val
		}
		result[sSplit[0]] = summmary
	}
	for k, v := range result {
		fmt.Println(k, v.ValSum, v.ValCount, v.ValMin, v.ValMax)
	}
	fmt.Println("Elapsed", time.Since(now))
}

func goParallelAlg() {
	now := time.Now()
	f, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	result := [10]map[string]Summary{}
	var batch [10][]string
	var i int64
	for scanner.Scan() {
		batch[i%4] = append(batch[i%4], scanner.Text())
		i++
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		result[i] = map[string]Summary{}
		go func() {
			defer wg.Done()
			for _, line := range batch[i] {
				sSplit := strings.Split(line, ";")
				val, err := strconv.ParseFloat(sSplit[1], 64)
				if err != nil {
					log.Fatal(err)
				}
				summmary := result[i][sSplit[0]]
				summmary.ValSum += val
				summmary.ValCount++
				if val > summmary.ValMax {
					summmary.ValMax = val
				}
				if val < summmary.ValMin {
					summmary.ValMin = val
				}
				result[i][sSplit[0]] = summmary
			}
		}()
	}
	wg.Wait()
	for _, vList := range result {
		for k, v := range vList {
			fmt.Println(k, v.ValSum, v.ValCount, v.ValMin, v.ValMax)
		}
	}
	fmt.Println("Elapsed", time.Since(now))
}

func concurrentWorker(result *[10]map[string]Summary, i int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for s := range ch {
		sSplit := strings.Split(s, ";")
		val, err := strconv.ParseFloat(sSplit[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		summmary := result[i][sSplit[0]]
		summmary.ValSum += val
		summmary.ValCount++
		if val > summmary.ValMax {
			summmary.ValMax = val
		}
		if val < summmary.ValMin {
			summmary.ValMin = val
		}
		result[i][sSplit[0]] = summmary
	}
	fmt.Println(1)
}

func goConcurrentAlg() {
	now := time.Now()
	f, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	wg := new(sync.WaitGroup)
	result := [10]map[string]Summary{}
	ch := make(chan string, 1000)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		result[i] = map[string]Summary{}
		go concurrentWorker(&result, i, wg, ch)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
	wg.Wait()
	for _, vList := range result {
		for k, v := range vList {
			fmt.Println(k, v.ValSum, v.ValCount, v.ValMin, v.ValMax)
		}
	}
	fmt.Println("Elapsed", time.Since(now))
}
