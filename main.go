package main

import (
	"fmt"
	"time"

	"github.com/wangbuer1984/myweb3/task1"
	"github.com/wangbuer1984/myweb3/task2"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := task1.SingleNumber(nums)
	fmt.Println("The single number is:", result)

	isPalindrome := task1.IsPalindrome(10)
	fmt.Println("Is 10 a palindrome?", isPalindrome)
	p := 10
	p2 := task2.MyPoint(&p)
	fmt.Println("task2.MyPoint?", p2)

	mySliec := []int{1, 2, 3, 4, 5}
	task2.MySlicePoint(&mySliec)
	fmt.Println(mySliec)

	tasks := []task2.Task{
		func() {
			fmt.Println("任务1开始")
			time.Sleep(500 * time.Microsecond)
			fmt.Println("任务1完成")
		},
		func() {
			fmt.Println("任务2开始")
			time.Sleep(200 * time.Microsecond)
			fmt.Println("任务2完成")
		},
		func() {
			fmt.Println("任务3开始")
			time.Sleep(400 * time.Millisecond)
			// 模拟任务失败
			panic("任务3发生意外错误")
		},
	}
	scheduler := task2.NewScheduler(tasks)
	fmt.Println("任务开始")
	scheduler.Run()
	// 打印执行结果
	scheduler.PrintResults()

	rectanglee := task2.Rectangle{}
	circle := task2.Circle{}
	rectanglee.Area()
	circle.Area()

	emp := task2.Employee{
		EmployeeID: 2,
		Person: task2.Person{
			Name: "wangdezhen",
			Age:  41,
		},
	}
	ss := emp.PrintInfo()
	fmt.Println(ss)

	// task2.MyBufferChain()
	task2.MyCounter()
	// task2.MyAtomicCounter()
}

// go func(id int) {
// 			defer wg.Done()
// 			for j:=0;j<1000;j++{
// 				mu.Lock()
// 				counter ++
// 				mu.Unlock()
// 			}
// 			fmt.Printf(协程%d 完成",id)
// 		}(i)
