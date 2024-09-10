package main

import (
	"MtCI-DSA/arrays"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {

	// ex2()
	// usingArrays()
	// fmt.Println(reverseAString("Cleverics"))

	// left := []int{4, 5, 67, 776}
	// right := []int{14, 67, 68, 408}

	// fmt.Println(mergeSortedArrays(left, right))

	// nums := []int{3, 2, 4}

	// fmt.Println(twoSum(nums, 6))

	testData := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
		{-2, 1},
		{-2, -1},
	}

	for _, input := range testData {
		// fmt.Println(maxSubArrayByKadane(input))
		fmt.Println(maxSubArrayDC(input, 0, len(input)-1))
	}

}

// finds the continuous part of the array, consisting of the elements that give max value
// when summed up, returns this very value
// using divide and conquer technique
// Returns sum of maximum sum subarray in nums[leftPoint ... rightPoint]
func maxSubArrayDC(nums []int, leftPoint, rightPoint int) int {

	// error handling: range like that can't exist
	if leftPoint > rightPoint {
		return math.MinInt
	}

	// the only element
	if leftPoint == rightPoint {
		return nums[leftPoint]
	}

	// Calculate the mid-point
	midPoint := (leftPoint + rightPoint) / 2

	// Calculate the maximum value:
	// max value is in the leftPart of nums
	// max value is in the rightPart of nums
	// max value is in the intersection of right and left parts

	return maxIntOfThree(
		maxSubArrayDC(nums, leftPoint, midPoint-leftPoint),
		maxSubArrayDC(nums, midPoint+1, rightPoint),
		maxIntersectionArraySum(nums, leftPoint, midPoint, rightPoint),
	)

}

//maxSubArrayDCHelpers{

// find the maximum possible sum so subarray of m would be its part
func maxIntersectionArraySum(nums []int, leftPoint int, middlePoint int, rightPoint int) int {

	// Grab the elements to the left of the mid-element [leftPoint, middlePoint]
	sum := 0
	leftPartSum := math.MinInt

	for i := middlePoint; i >= leftPoint; i-- {
		sum += nums[i]
		if sum > leftPartSum {
			leftPartSum = sum
		}
	}

	// Grab the elements to the right of the mid-element [middlePoint, rightPoint]
	sum = 0
	rightPartSum := math.MinInt
	for i := middlePoint; i <= rightPoint; i++ {
		sum += nums[i]
		if sum > rightPartSum {
			rightPartSum = sum
		}
	}

	// Return the sum of elements on the left and on the right of the middle point,

	return maxIntOfThree(leftPartSum+rightPartSum-nums[middlePoint], leftPartSum, rightPartSum)

}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxIntOfThree(a, b, c int) int {
	subValue := maxInt(a, b)
	return maxInt(subValue, c)
}

//}maxSubArrayDCHelpers

// finds the continuous part of the array, consisting of the elements that give max value
// when summed up, returns this very value
// Kadane's algorithm: O(n^2)
func maxSubArrayByKadane(nums []int) int {

	result := nums[0]

	// starting point of  the subarray
	for index, _ := range nums {

		currentSum := 0

		// ending point of the subarray

		for j := index; j < len(nums); j++ {
			currentSum += nums[j]

			// update result if currentSum overlaps it
			if currentSum > result {
				result = currentSum
			}
		}

	}

	return result

}

// returns the indices of the elements equal to the pair of summands
func twoSum(nums []int, target int) []int {

	var result []int

	for i, v := range nums {

		nextValue := target - v
		fmt.Println(nums[i+1:])
		secondOperandPosition := slices.Index(nums[i+1:], nextValue)

		fmt.Println(v)
		fmt.Println(nextValue)
		fmt.Println(secondOperandPosition)

		if secondOperandPosition == -1 {
			continue
		} else {
			result = append(result, i)
			result = append(result, i+secondOperandPosition+1)
			break
		}

	}

	return result

}

func mergeSortedArrays(left []int, right []int) []int {

	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	i := 0
	j := 0

	var result []int

	for k := 0; k < len(right)+len(left)-1; k++ {

		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}

	}

	return result

}

func reverseAString(input string) string {

	runes := []rune(input)

	var sb strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}

	return sb.String()
}

func usingArrays() {

	arrCapacity := 9
	arr := arrays.New(arrCapacity)

	// testing Push
	arr.Push(10)

	value, err := arr.Get(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// testing Pop
	err = arr.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	for i := 0; i < 3; i++ {
		arr.Push(i)
	}
	fmt.Println(arr)

	// testing deletion
	arr.Delete(1)
	fmt.Println(arr)

}

func ex2() {

	var collection = []int{6, 4, 3, 2, 1, 7}
	var sum = 9
	// this approach sucks, as the time complexity is O(n^2) and it could be improved if some more memory is sacrificed
	naiveHasPairWithSum(collection, sum)

	// reimplemented version using data structures, has better time complexity
	betterHasPairWithSum(collection, sum)

}

func naiveHasPairWithSum(collection []int, sum int) bool {

	// get the collection length
	len := len(collection)

	// if there are no values return false
	if len == 0 {
		return false
	}

	// iterate over each element of the array

	// check if the given element forms the sum with any of the others
	for cnt := 0; cnt < len-1; cnt++ {
		for nCnt := cnt + 1; nCnt < len; nCnt++ {
			// if it does terminate the code and return true
			if collection[cnt]+collection[nCnt] == sum {
				return true
			}
		}
	}
	//if the result is negative but the traversal is over, return false
	return false

}

func betterHasPairWithSum(collection []int, sum int) bool {

	// if the collection is empty - quit
	if len(collection) == 0 {
		return false
	}

	// use a map for a quick (O(1)) lookup
	var mappo = make(map[int]bool)

	// traverse the collection
	for _, v := range collection {

		// calculate the second operand as sum-op1 and check if the next checked value is equal to it

		ok, _ := mappo[v]
		if ok {
			return true
		}

		mappo[sum-v] = true
	}

	// if no value is found, return false
	return false

}

func ex1() {

	var first = []rune{'a', 'b', 'c', 'x'}
	var last = []rune{'z', 'y', 'i'}

	var last_but_one = []rune{'z', 'y', 'i'}
	findCommonItems3N(first, last)
	findCommonItemsN2(first, last_but_one)

}

func findCommonItemsN2(first []rune, second []rune) bool {

	// traverse over each element of the first array
	for _, v := range first {

		// traverse over each el of the second array

		for _, nv := range second {

			// if found => terminate the code
			if v == nv {
				return true
			}

		}

	}

	return false
}

func findCommonItems3N(first []rune, second []rune) bool {

	var mappo = make(map[rune]int)
	var mappie = make(map[rune]int)

	for k, v := range first {
		mappo[v] = k
	}

	for k, v := range second {
		mappie[v] = k
	}

	traversedCollection := mappie
	anotherOne := mappo
	if len(mappo) < len(mappie) {
		traversedCollection = mappo
		anotherOne = mappie
	}

	for k, _ := range traversedCollection {

		_, ok := anotherOne[k]
		if ok {
			return ok
		}

	}

	return false

}

func ex0() {
	boxes := [5]int{1, 2, 3, 4, 5}

	result := ""

	for _, v := range boxes {

		for _, n := range boxes {
			result = fmt.Sprintf("%s\n (%d : %d)", result, v, n)
		}

	}

	fmt.Print(result)

}
