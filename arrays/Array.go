package arrays

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"slices"
	"strings"
)

type Array struct {
	length   int
	capacity int
	data     []interface{}
}

func New(capacity int) *Array {
	return &Array{
		length:   0,
		capacity: capacity,
		data:     []any{},
	}
}

func (a *Array) Get(index int) (any, error) {
	if len(a.data) == 0 || len(a.data) < index {
		return nil, errors.New("index out of range bounds")
	} else {
		return a.data[index], nil
	}
}

func (a *Array) Push(item any) error {

	if len(a.data) > 0 && (reflect.TypeOf(item) != reflect.TypeOf(a.data[0])) {
		return errors.New("failed to add value of type " + reflect.TypeOf(item).String())
	}

	if a.capacity < a.length {
		a.capacity *= 2
	}

	a.data = append(a.data, item)
	a.length++

	return nil

}

func (a *Array) Pop() error {

	if a.length == 0 {
		return errors.New("failed to pop an item from an empty collection")
	} else {

		a.data = (a.data[:len(a.data)-1])
		a.length--
		return nil

	}

}

func (a *Array) Delete(index int) error {

	if a.length == 0 || index > a.length-1 {
		return errors.New("index out of range bounds")
	} else if index == a.length-1 {
		a.Pop()
	} else {

		a.data = append(a.data[:index], a.data[index+1:]...)
		a.length--

	}

	return nil
}

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]

func RotateLikeABoss(nums []int, k int) {

	if k == 0 {
		return
	}

	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}

func reverse(nums []int, start int, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}

func RotateArray(nums []int, k int) {

	if len(nums) < 2 || k < 1 {
		return
	}

	for cnt := 1; cnt <= k; cnt++ {

		shiftedNums := make([]int, len(nums))
		for i := range nums {

			nextI := i + 1
			if nextI >= len(nums) {
				nextI -= len(nums)
			}

			shiftedNums[nextI] = nums[i]
		}

		for k, v := range shiftedNums {
			nums[k] = v
		}
	}

}

func ContainsDuplicates(nums []int) bool {

	mappo := make(map[int]int)
	for _, v := range nums {

		_, found := mappo[v]

		if found {
			return true
		} else {
			mappo[v]++
		}
	}

	return false

}

func MoveZeroesInTheArray(nums []int) {

	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}

}

func MoveZeroes(input []int) []int {

	result := []int{}
	if len(input) < 2 {
		return input
	}

	// sort the array
	input = mergeSort(input)

	// calculate the shift value  (the zero-block length)
	firstOccurence := slices.Index(input, 0)
	currentPosition := firstOccurence

	if currentPosition > -1 {

		for {
			if input[currentPosition] != 0 {
				break
			} else {
				currentPosition++
			}
		}

	}

	// swap zeroes
	if firstOccurence > 0 {
		result = append(result, input[:firstOccurence-1]...)
	}
	result = append(result, input[currentPosition:]...)

	nonZeroBlockLength := len(result)

	for i := nonZeroBlockLength; i < len(input); i++ {
		result = append(result, 0)
	}

	return result
}

//moveZeroes helpers{

func merge(leftPart, rightPart []int) []int {

	result := make([]int, 0, len(leftPart)+len(rightPart))

	i, j := 0, 0

	for i < len(leftPart) && j < len(rightPart) {

		if leftPart[i] < rightPart[j] {
			result = append(result, leftPart[i])
			i++
		} else {
			result = append(result, rightPart[j])
			j++
		}

	}

	result = append(result, leftPart[i:]...)
	result = append(result, rightPart[j:]...)

	return result

}

func mergeSort(input []int) []int {

	if len(input) < 2 {
		return input
	}

	middlePoint := len(input) / 2
	leftPart := input[:middlePoint]
	rightPart := input[middlePoint:]

	sortedLeft := mergeSort(leftPart)
	sortedRight := mergeSort(rightPart)

	return merge(sortedLeft, sortedRight)

}

// }moveZeroes helpers

func KickAssMaxSubArray(nums []int) int {

	var maxSubsum = nums[0]
	var currentSum = 0

	for _, num := range nums {

		if currentSum < 0 {
			currentSum = 0
		}

		currentSum += num
		maxSubsum = max(maxSubsum, currentSum)

	}

	return maxSubsum

}

// finds the continuous part of the array, consisting of the elements that give max value
// when summed up, returns this very value
// using divide and conquer technique
// Returns sum of maximum sum subarray in nums[leftPoint ... rightPoint]
func MaxSubArrayDC(nums []int, leftPoint, rightPoint int) int {

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
		MaxSubArrayDC(nums, leftPoint, midPoint-leftPoint),
		MaxSubArrayDC(nums, midPoint+1, rightPoint),
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
func MaxSubArrayByKadane(nums []int) int {

	result := nums[0]

	// starting point of  the subarray
	for index := range nums {

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
	arr := New(arrCapacity)

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
	FindCommonItems3N(first, last)
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

func FindCommonItems3N(first []rune, second []rune) bool {

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
