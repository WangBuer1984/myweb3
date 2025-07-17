package task2

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func MyPoint(p *int) int {
	p1 := *p
	p1 += 10
	return p1
}

func MySlicePoint(slicePtr *[]int) {
	tmpslicePtr := *slicePtr
	for i := range tmpslicePtr {
		tmpslicePtr[i] = tmpslicePtr[i] * 2

	}

}

// 打印奇偶数
func MyGoroutine() {

	var wg sync.WaitGroup
	wg.Add(2)

	//奇数打印
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数：", i)
		}
	}()

	wg.Wait()
	fmt.Println("完成")
}

//并发任务调度

type Task func()

type TaskResult struct {
	ID       int
	Duration time.Duration
	Error    error
}

type Scheduler struct {
	tasks   []Task
	results []TaskResult
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func NewScheduler(tasks []Task) *Scheduler {
	return &Scheduler{
		tasks:   tasks,
		results: make([]TaskResult, len(tasks)),
	}
}

// executeTask 执行单个任务并记录结果

func (s *Scheduler) executeTask(id int, task Task) {
	defer s.wg.Done()
	start := time.Now()

	defer func() {
		if r := recover(); r != nil {
			s.mu.Lock()
			s.results[id] = TaskResult{
				ID:       id,
				Duration: time.Since(start),
				Error:    fmt.Errorf("panic:%v", r),
			}
			s.mu.Unlock()
		}
	}()
	task()
	s.mu.Lock()
	s.results[id] = TaskResult{
		ID:       id,
		Duration: time.Since(start),
	}

	s.mu.Unlock()

}

// Run 并发执行所有任务
func (s *Scheduler) Run() {
	s.wg.Add(len(s.tasks))
	for i, task := range s.tasks {
		go s.executeTask(i, task)
	}
	s.wg.Wait()
}

// PrintResults 打印任务执行结果
func (s *Scheduler) PrintResults() {
	fmt.Println("\n任务执行结果:")
	fmt.Printf("%-6s %-12s %-10s\n", "任务ID", "耗时", "状态")

	for _, res := range s.results {
		status := "成功"
		if res.Error != nil {
			status = fmt.Sprintf("失败: %v", res.Error)
		}
		fmt.Printf("%-6d %-12s %-10s\n",
			res.ID,
			res.Duration.Round(time.Millisecond),
			status)
	}
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle.Area")
}

func (c Circle) Area() {
	fmt.Println("Circle.Area")
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() string {
	return "name:" + e.Name + "age:" + strconv.Itoa(e.Age) + "id=" + strconv.Itoa(e.EmployeeID)
	//
}

func MyChain() {
	intChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(intChan)
		fmt.Println("生产者开始")
		for i := 0; i < 10; i++ {
			fmt.Printf("发送：%d\n", i)
			intChan <- i
		}
		fmt.Println("生产者发送完毕")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("消费者接收：")
		for num := range intChan {
			fmt.Printf("接收：%d\n", num)
		}
		fmt.Println("消费者接收完毕")
	}()

	wg.Wait()
	fmt.Println("done")
}
func MyBufferChain() {
	intChan := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(intChan)
		fmt.Println("生产者开始")
		for i := 0; i < 100; i++ {
			fmt.Printf("发送：%d\n", i)
			intChan <- i
		}
		fmt.Println("生产者发送完毕")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("消费者接收：")
		for num := range intChan {
			fmt.Printf("接收：%d\n", num)
		}
		fmt.Println("消费者接收完毕")
	}()

	wg.Wait()
	fmt.Println("done")
}
func MyCounter() {
	var counter int
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("协程%d启动\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("最终计数器值为：", counter)
}
func MyAtomicCounter() {
	var counter int64

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)

			}
			fmt.Printf("协程%d启动\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("最终计数器值为：", counter)
}
