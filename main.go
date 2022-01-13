package main

import (
	"fmt"
	"math"
)

func main() {
	requirements := []string{
		"gym", "school", "store",
	}

	testBlocks := []map[string]bool{
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    true,
			"school": false,
			"store":  false,
		},
		{
			"gym":    true,
			"school": true,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  true,
		},
	}

	fmt.Println(bestBlockIndex(requirements, testBlocks))
}

func bestBlockIndex(req []string, blocks []map[string]bool) (result int) {
	rowLength := len(req) + 1
	blocksLength := len(blocks)
	dp := make([][]int, blocksLength)

	for i := 0; i < blocksLength; i++ {
		dp[i] = make([]int, rowLength)
		for ii := 0; ii < rowLength; ii++ {
			dp[i][ii] = 100
		}
	}

	// iterate first block
	for i, building := range req {
		if blocks[0][building] {
			dp[0][i] = 0
		}
	}

	printMatrix(dp)

	// iterate other blocks
	for i := 1; i < blocksLength; i++ {
		for j, v := range req {
			if blocks[i][v] {
				dp[i][j] = 0
			} else {
				if dp[i-1][j] != 100 {
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j]+1)))
				}
			}
			dp[i][rowLength-1] = int(math.Max(float64(dp[i][rowLength-1]), float64(dp[i][j])))
		}
	}

	printMatrix(dp)

	// iterate backwards
	for i := blocksLength - 2; i > 0; i-- {
		for j, v := range req {
			if blocks[i][v] {
				dp[i][j] = 0
			} else {
				if dp[i+1][j] != 100 {
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i+1][j]+1)))
				}
			}
			dp[i][rowLength-1] = int(math.Max(float64(dp[i][rowLength-1]), float64(dp[i][j])))

			if dp[i][j] == 100 || dp[i][rowLength-1] == 100 {
				dp[i][rowLength-1] = -1
				dp[i][j] = -1
			}

			result = int(math.Min(float64(result), float64(dp[i][rowLength-1])))
		}
	}

	printMatrix(dp)

	return result
}

func printMatrix(dp [][]int) {
	fmt.Println()
	for i := 0; i < len(dp); i++ {
		for ii := 0; ii < len(dp[0]); ii++ {
			fmt.Printf(" %v", dp[i][ii])
		}
		fmt.Println()
	}
	fmt.Println()
}
