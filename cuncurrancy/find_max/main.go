package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var ArraySizePerGoroutine = 10000
	var LengthOfArry = 10000000

	d := float64(LengthOfArry) / float64(ArraySizePerGoroutine)
	NumberOfRoutines := int(math.Ceil(d))

	fmt.Println("Number of go routines=", NumberOfRoutines)

	var NumberArray = make([]int, LengthOfArry)

	for index := range NumberArray {
		NumberArray[index] = rand.Intn(1000)
	}

	//for testing
	NumberArray[5000] = 88888

	fmt.Println("Making of array is done")

	start := time.Now()

	queue := make(chan int, NumberOfRoutines)
	var wg sync.WaitGroup

	for i := 0; i < NumberOfRoutines; i++ {
		var start, end int
		start = i * ArraySizePerGoroutine
		if LengthOfArry-start < ArraySizePerGoroutine {
			end = LengthOfArry
		} else {
			end = (i + 1) * ArraySizePerGoroutine
		}
		sub := NumberArray[start:end]

		wg.Add(1)
		n := i + 1

		go func(subarray []int, chn chan int, rn int) {
			chn <- Findmax(subarray)
			//fmt.Println("goroutiene ", rn, " ended")
			defer wg.Done()
		}(sub, queue, n)
		//fmt.Println("goroutiene ", n, " started")
	}

	wg.Wait()
	close(queue)
	var finalSlice []int
	for elem := range queue {
		finalSlice = append(finalSlice, elem)
	}

	fmt.Println("Max =", Findmax(finalSlice))

	elapsed := time.Since(start)
	log.Printf("Find max took %s", elapsed)
}

//Findmax finds the max in a slice of int
func Findmax(ArraySlice []int) int {
	var max = 0
	for i, v := range ArraySlice {
		if ArraySlice[max] < v {
			max = i
		}
	}
	return ArraySlice[max]
}
