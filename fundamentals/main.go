package main

import (
	"fmt"
	"math"
)

func main() {

	// second greatest value
	vals := []int{335, 3, 2, 5, 6, 8, 11}
	result := secondGreatest(vals)
	fmt.Println(fmt.Sprintf("input: %v, result: %d", vals, result))

	/*	Get the product of the following two 64-digit dec numbers:
		3141592653589793238462643383279502884197169399375105820974944592 and
		2718281828459045235360287471352662497757247093699959574966967627
		using Karatsuba's multiplication algorithm */

	// a := int256.NewInt(3141592653589793238462643383279502884197169399375105820974944592)
	// b := int256.NewInt(2718281828459045235360287471352662497757247093699959574966967627)
	// a := big.NewInt(4565)
	// b := big.NewInt(64481)

	a, b := int64(1492), int64(1147)
	bigResult := multiplyKaratsubaWay(a, b)

	fmt.Printf("%d * %d  = %d\n", a, b, bigResult)

}

/*
	Implement Karatsuba's integer multiplication algorithm. To get the most out of this problem

the code should invoke multiplication operator only on pairs of single-digit numbers.
Get the product of the following two 64-digit dec numbers:
3141592653589793238462643383279502884197169399375105820974944592 and
2718281828459045235360287471352662497757247093699959574966967627
*/

func multiplyKaratsubaWay(a int64, b int64) int64 {

	var maxSize uint
	isPositive := true

	if a == 0 || b == 0 {
		return 0
	}

	isPositive = !(a > 0 && b < 0) || (a < 0 && b > 0)

	if a < 0 {
		a *= -1
	}

	if b < 0 {
		b *= -1
	}

	if a < 10 || b < 10 {
		return a * b
	}

	aSize := getNumberSize(a)
	bSize := getNumberSize(b)

	if aSize >= bSize {
		maxSize = aSize / 2
	} else {
		maxSize = bSize / 2
	}

	aLeftPart, aRightPart := getDigits(a, maxSize)
	bLeftPart, bRightPart := getDigits(b, maxSize)

	firstSubpart := multiplyKaratsubaWay(int64(aRightPart), int64(bRightPart))

	middleAElement := aLeftPart + aRightPart
	middleBElement := bLeftPart + bRightPart

	middleSubpart := multiplyKaratsubaWay(middleAElement, middleBElement)

	finalSubpart := multiplyKaratsubaWay(aLeftPart, bLeftPart)

	result := finalSubpart*int64(math.Pow(10, float64(2*maxSize))) + (middleSubpart-finalSubpart-firstSubpart)*int64(math.Pow(10, float64(maxSize))) + firstSubpart
	if !isPositive {
		result *= -1
	}

	return result

}

func getNumberSize(number int64) uint {

	var result uint

	if number == 0 {
		return 1
	} else if number < 0 {
		number = -number
	}

	for number > 0 {
		result++
		number /= 10
	}

	return result

}

func getDigits(number int64, digits uint) (int64, int64) {

	divisor := int64(math.Pow(10, float64(digits)))

	if number >= divisor {
		return number / divisor, number % divisor
	} else {
		return 0, number
	}

}

/* Given an input of an unsorted array of n distinct numbers, where n is a power of 2.
Give an algorithm that identifies the second-largest number in the array and that uses
at most n + log2(N) -2. (Hint: Consider computing the largest number using a knockout
tournament. )*/

func secondGreatest(input []int) int {

	result := 0

	// implement a mergesort for log2(N)
	sortedResult := mergeSort(input)

	//grab the item len(result) - 2, where len(result)-1 is the index of the last position
	position := len(input) - 2
	result = sortedResult[position]

	return result

}

func mergeSort(input []int) []int {

	inputLen := len(input)
	if inputLen < 2 {
		return input
	}

	leftPart := mergeSort(input[:inputLen/2])
	rightPart := mergeSort(input[inputLen/2:])

	return merge(leftPart, rightPart)

}

func merge(leftPart []int, rightPart []int) []int {

	i := 0
	j := 0

	// resultSize := len(leftPart) + len(rightPart)
	result := []int{}

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
