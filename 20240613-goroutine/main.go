package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Value struct {
	Station string
	Sum     float32
	// uint32 max value is 4,294,967,295, we have 100,000,000 rows
	Count uint32
	Max   float32
	Min   float32
}

func splitLine(line string) (LineSplit, error) {
	// we don't use csv parser since no special characters exists to save computation power
	lineSplit := strings.SplitN(line, ";", 2)
	if len(lineSplit) < 2 {
		return LineSplit{}, fmt.Errorf("invalid line: %s", line)
	} else {
		return LineSplit{lineSplit[0], lineSplit[1]}, nil
	}
}

type LineSplit struct {
	Station     string
	Temperature string
}

func decodeLine(lineSplit LineSplit, result map[string]*Value) error {
	elm := result[lineSplit.Station]
	if elm == nil {
		elm = new(Value)
		// temperature value is between -99.9 .. 99.9
		elm.Max = -100
		elm.Min = 100
		elm.Station = lineSplit.Station
		result[lineSplit.Station] = elm
	}
	value, err := strconv.ParseFloat(lineSplit.Temperature, 32)
	if err != nil {
		return err
	}
	value32 := float32(value)
	elm.Sum += value32
	elm.Count += 1
	if value32 > elm.Max {
		elm.Max = value32
	}
	if value32 < elm.Min {
		elm.Min = value32
	}
	return nil
}

func SplitWorker(wg *sync.WaitGroup, lineChannel <-chan string, decodeChannel chan<- LineSplit) {
	defer wg.Done()
	for line := range lineChannel {
		lineSplit, err := splitLine(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		decodeChannel <- lineSplit
	}
}

func DecodeWorker(wg *sync.WaitGroup, decodeChannel <-chan LineSplit, result map[string]*Value) {
	defer wg.Done()
	for line := range decodeChannel {
		err := decodeLine(line, result)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	start := time.Now()
	f, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 64k buffer by default
	// in the github page, max line will be 100 bytes for station, 1 byte separator, and 5 bytes temperature
	const maxCapacity = 128 // round capacity to 128 bytes per line
	scanner := bufio.NewScanner(f)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	decodeWg := new(sync.WaitGroup)
	splitWg := new(sync.WaitGroup)
	lineChannel := make(chan string, 100)
	decodeChannel := make(chan LineSplit, 100)
	result := make(map[string]*Value)
	var workerResult []map[string]*Value
	for i := 0; i < runtime.NumCPU(); i++ {
		splitWg.Add(1)
		go SplitWorker(splitWg, lineChannel, decodeChannel)
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		decodeWg.Add(1)
		decodeResult := make(map[string]*Value)
		workerResult = append(workerResult, decodeResult)
		go DecodeWorker(decodeWg, decodeChannel, decodeResult)
	}

	// room for optimization
	for scanner.Scan() {
		lineChannel <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	close(lineChannel)
	splitWg.Wait()
	close(decodeChannel)
	decodeWg.Wait()
	processTime := time.Now()

	for _, decodeResult := range workerResult {
		for k, v := range decodeResult {
			if resultV := result[k]; resultV == nil {
				result[k] = v
			} else {
				resultV.Sum += v.Sum
				resultV.Count += v.Count
				if v.Max > resultV.Max {
					resultV.Max = v.Max
				}
				if v.Min < resultV.Min {
					resultV.Min = v.Min
				}
			}
		}
	}
	var resultList []*Value
	for _, value := range result {
		resultList = append(resultList, value)
	}
	mergeTime := time.Now()
	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].Station < resultList[j].Station
	})
	sortTime := time.Now()
	output, err := os.Create("output.csv")
	defer output.Close()
	if err != nil {
		panic(err)
	}
	for _, value := range resultList {
		_, err = output.WriteString(fmt.Sprintf("%s;%f;%f;%f\n", value.Station, value.Sum/float32(value.Count), value.Min, value.Max))
		if err != nil {
			panic(err)
		}
	}
	outputTime := time.Now()
	fmt.Println(fmt.Sprintf("processing time: %.3f", processTime.Sub(start).Seconds()))
	fmt.Println(fmt.Sprintf("merge time: %.3f", mergeTime.Sub(processTime).Seconds()))
	fmt.Println(fmt.Sprintf("sort time: %.3f", sortTime.Sub(mergeTime).Seconds()))
	fmt.Println(fmt.Sprintf("dump time: %.3f", outputTime.Sub(sortTime).Seconds()))
	fmt.Println(fmt.Sprintf("total time: %.3f", outputTime.Sub(start).Seconds()))
}
