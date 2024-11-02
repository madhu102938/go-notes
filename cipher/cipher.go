package main

import "fmt"

func create2DSlice[T comparable](n, m int, x T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, m)
		for j := range dp[i] {
			dp[i][j] = x
		}
	}

	return dp
}

func max[T int | float32 | float64] (a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

func actualAnswer(i, j int, grid, dp [][]int) int {
	if (i == len(grid)-1) {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	idx := []int{-1, 1, 0, 0}
	idy := []int{0, 0, -1, 1}
	ans, moves := 0, 0

	for k := 0; k < 4; k++ {
		newi, newj := i + idx[k], j + idy[k]
		if (newi >= 0 && newj >= 0 && newi < len(grid) && newj < len(grid[0]) && grid[i][j] < grid[newi][newj]) {
			moves = max(actualAnswer(newi, newj, grid, dp) + 1, moves)
			
			if (newi == len(grid[0])-1) {
				ans = moves
			}

			// ans = max(ans, moves)
			// if ans != 0 {
			// 	fmt.Println(ans, i, j, newi, newj, grid[i][j], grid[newi][newj])
			// }
		}
	}

	dp[i][j] = ans
	return ans
}

func maxMoves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := create2DSlice(n, m, -1)

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, actualAnswer(i, 0, grid, dp))
	}
	
	return ans
}

func main() {
	grid := [][]int{{3,2,4},{2,1,9},{1,1,7}}
	fmt.Println(maxMoves(grid))
}