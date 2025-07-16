package main

import (
	"fmt"

	"github.com/wangbuer1984/myweb3/task1"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := task1.SingleNumber(nums)
	fmt.Println("The single number is:", result)

	isPalindrome := task1.IsPalindrome(10)
	fmt.Println("Is 10 a palindrome?", isPalindrome)
}
