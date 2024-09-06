package main

import "fmt"

func main() {

	vals := []int{335, 3, 2, 5, 6, 8, 11}
	result := secondGreatest(vals)
	fmt.Println(fmt.Sprintf("input: %v, result: %d", vals, result))

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

//

/* Implement Karatsuba's integer multiplication algorithm. To get the most out of this problem
the code should invoke multiplication operator only on pairs of single-digit numbers.
Get the product of the following two 64-digit dec numbers:
3141592653589793238462643383279502884197169399375105820974944592 and
2718281828459045235360287471352662497757247093699959574966967627*/
