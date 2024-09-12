package main

import (
	"MtCI-DSA/arrays"
	"fmt"
)

func main() {

	// arraySession()
	hashMapSession()
}

func hashMapSession() {

}

func arraySession() {

	testArrayData := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
		{-2, 1},
		{-2, -1},
		{0, -3, -2, -3, -2, 2, -3, 0, 1, -1},
		{0, 1, 0, 3, 12},
		{0},
		{1, 2, 3, 4, 5, 6, 7},
	}

	for _, input := range testArrayData {

		fmt.Println(arrays.MaxSubArrayByKadane(input))
		fmt.Println(arrays.MaxSubArrayDC(input, 0, len(input)-1))
		fmt.Println(arrays.KickAssMaxSubArray(input))

		fmt.Println(arrays.MoveZeroes(input))

		arrays.MoveZeroesInTheArray(input)
		fmt.Println(input)

		fmt.Println(arrays.ContainsDuplicates(input))

		arrays.RotateArray(input, 3)
		arrays.RotateLikeABoss(input, 3)
		fmt.Print(input)

	}

}
