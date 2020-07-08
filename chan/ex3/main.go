package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	Value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func z(zl chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			Value: x,
		}
		zl <- newJob
		time.Sleep(time.Second)
	}
}

func b(zl <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-zl
		sum := int64(0)
		n := job.Value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go z(jobChan)
	go b(jobChan, resultChan)

	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.Value, result.sum)
	}

	wg.Wait()
}
