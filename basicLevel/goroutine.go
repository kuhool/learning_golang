package basicLevel

import (
	"math/rand"

	"fmt"
	"runtime"
	"sync"
	"time"
)

// init的使用方法？？？
func init() {
	fmt.Println("Goroutine lets go ")
}

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	Job *Job
	Sum int
}

// worker pool 实现一个协程池
func GoroutineT() {

	//定义一个消费管道
	JobChan := make(chan *Job, 128)
	//定义一个结果输出管道
	ResultChan := make(chan *Result, 128)
	var num int = 64
	createPool(num, JobChan, ResultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan { //打印
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.Job.Id, result.Job.RandNum, result.Sum)
		}
	}(ResultChan)

	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		JobChan <- job
	}

}

// 创建协程池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {

	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				//随机数
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				r := &Result{
					Job: job,
					Sum: sum,
				}

				//运算结果扔到管道
				resultChan <- r
			}

		}(jobChan, resultChan)
	}

}

func GoroutineT5() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 9; i++ {
		go func(num int) {
			for {
				fmt.Println("CPU-intensive task...", num)
				time.Sleep(time.Millisecond * 1000)
				runtime.Gosched() // 让出CPU时间片
			}
		}(i)
	}

	time.Sleep(time.Second * 1)
}
func GoroutineT4() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine1:", i)
			runtime.Gosched()
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("main goroutine:", i)
	}

}

func GoroutineT2() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		fmt.Println("hello")
		runtime.Gosched()
	}

}
func GoroutineT1() {
	// 合起来写
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println("Hello Goroutine!!!!!", num)
		}(i)
	}

	wg.Wait()
}

func Goroutine3() {

	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
