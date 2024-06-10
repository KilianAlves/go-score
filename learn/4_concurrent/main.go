package main

import (
	"fmt"
	"math"
	"time"
)

func pi(iterStart uint, iterEnd uint) float64 {
	var result float64
	for i := float64(iterStart); i < float64(iterEnd); i++ {
		result += math.Pow(-1, i) / (2*i + 1)
	}
	return 4 * result
}

func piRoutine(iterStart uint, iterEnd uint, ch chan float64) {
	pi := pi(iterStart, iterEnd)
	ch <- pi
}

func piConcurrent(iter uint, taskCount uint) float64 {
	ch := make(chan float64, taskCount)

	iterationsPerTask := iter / taskCount

	for i := uint(0); i < taskCount; i++ {
		start := i * iterationsPerTask
		end := start + iterationsPerTask
		if i == taskCount-1 {
			end = iter
		}
		go piRoutine(start, end, ch)
	}

	piResult := 0.0
	for i := uint(0); i < taskCount; i++ {
		piResult += <-ch
	}

	return piResult
}

func main() {

	start1 := time.Now()
	pi := pi(0, 1200000000)
	end1 := time.Now()

	start2 := time.Now()
	piConcurrent2 := piConcurrent(1200000000, 2)
	end2 := time.Now()

	start3 := time.Now()
	piConcurrent3 := piConcurrent(1200000000, 4)
	end3 := time.Now()

	fmt.Println("1 thread :", pi, end1.Sub(start1))
	fmt.Println("2 thread :", piConcurrent2, end2.Sub(start2))
	fmt.Println("plusieurs thread :", piConcurrent3, end3.Sub(start3))
}
