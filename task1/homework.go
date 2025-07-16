package task1

import (
	"sort"
	"strconv"
)

// 只出现一次的数字
func SingleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

// 回文数
func IsPalindrome(x int) bool {
	if x <= 0 {
		return false
	}

	s := strconv.Itoa(x)

	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--

	}
	return true
}

// 有效的括号
func IsValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if open, exists := mapping[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0

}

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			break
		}
	}
	prefix := ""
	for i := 0; i < minLen; i++ {
		currentChar := strs[0][i]
		for _, str := range strs {
			if str[i] != currentChar {
				return prefix
			}
		}
		prefix += string(currentChar)
	}
	return prefix
}

//. 加一

func PlusOne(digits []int) []int {
	num := 0
	for _, d := range digits {
		num = num*10 + d
	}
	num++
	result := []int{}
	if num == 0 {
		return []int{0}
	}
	numStr := strconv.Itoa(num)
	for i, c := range numStr {
		result[i] = int(c - '0')
	}
	return result
}

// 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// 合并区间
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] <= last[1] { // 有重叠
			last[1] = max(last[1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

// 两数之和
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		complement := target - num
		if _, found := numMap[complement]; found {
			return []int{complement, num}
		}
		numMap[num] = true
	}
	return nil
}
