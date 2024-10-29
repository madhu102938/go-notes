package main

import (
    "sort"
    "fmt"
	// "math"
)

type myType []int

func (m myType) Less(i, j int) bool {
    return m[i] < m[j]
}

func (m myType) Len() int {
    return len(m)
}

func (m myType) Swap(i, j int) {
    m[j], m[i] = m[i], m[j]
}

func condition(num int) bool {
	return int64(num) * int64(num) > int64(1e5)
}

func longestSquareStreak(nums []int) int {
    sort.Sort(myType(nums))
	var ans int
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	
	mp := make(map[int]bool)
	for _, j := range nums {
		mp[j] = true
	}

	for _, num := range nums {

		temp := 1
		toFind := num
		for !condition(toFind) && mp[toFind * toFind] {
			temp++
			ans = max(temp, ans)
			toFind = toFind * toFind
		}
	}

	return ans
}

func main() {
	nums := []int{4,3,6,16,8,2}
	fmt.Println(longestSquareStreak(nums))
}