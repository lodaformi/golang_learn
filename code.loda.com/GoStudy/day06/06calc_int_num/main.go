package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	jobs *job
	sum  int64
}

//生成随机数，写入到job中
func randFunc(j chan<- *job) {
	defer wg.Done()
	for {
		tmp := &job{value: rand.Int63()}
		j <- tmp
		time.Sleep(time.Millisecond * 500)
	}
}

func calcSum(j <-chan *job, res chan<- *result) {
	defer wg.Done()
	for v := range j {
		num := v.value
		var s = int64(0)
		//计算每位数字之和
		for num > 0 {
			s += num % 10
			num /= 10
		}
		tmp := &result{jobs: v, sum: s}
		res <- tmp
	}
}

var wg sync.WaitGroup

func main() {
	//声明指针变量的channel
	jobChan := make(chan *job, 10)
	resultChan := make(chan *result, 10)

	wg.Add(1)
	go randFunc(jobChan)

	//开启24个goroutine计算
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go calcSum(jobChan, resultChan)
	}

	//输出结果
	for v := range resultChan {
		fmt.Println(v.jobs.value, v.sum)
	}
	wg.Wait()
}
