package main

import (
	"fmt"
	"indexer/index"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	totalTime := 0.0
	benchmarkResults := make([]time.Duration, 0, 10)
	for i := 1; i <= 5; i++ {
		start := time.Now()
		index.BuildIndex([]string{"data"}, 10000)
		elapsed := time.Since(start)
		benchmarkResults = append(benchmarkResults, elapsed)
		totalTime += float64(elapsed)/100_000_000
	}
	fmt.Println(benchmarkResults)
	fmt.Printf("Average execution time: %f \n", totalTime/10)
}
