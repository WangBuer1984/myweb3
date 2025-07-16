package main

import (
	"fmt"

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
}
