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

func piRoutine(iterStart uint, iterEnd uint, ch chan uint) {
	pi := pi(iterStart, iterEnd)
	ch <- uint(pi)
}

func piConcurrent(iter uint) uint {
	ch := make(chan uint)
	go piRoutine(0, iter, ch)
	res := <-ch
	go piRoutine(0, iter, ch)
	res += <-ch

	return res

}

func main() {
	start := time.Now()

	pi := pi(0, 1200000000)

	piConcurrent := piConcurrent(1200000000)

	end := time.Now()

	fmt.Println("1 thread :", pi, end.Sub(start))
	fmt.Println("2 thread :", piConcurrent, end.Sub(start))
}
