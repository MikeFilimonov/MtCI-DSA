package sorting

func BubbleSort(input []int) []int {

	for i := 0; i < len(input); i++ {

		for j := i + 1; j < len(input); j++ {

			if input[i] > input[j] {
				input[j], input[i] = input[i], input[j]
			}

		}

	}

	return input

}

func SelectionSort(input []int) []int {

	for i := 0; i < len(input); i++ {

		minimum := i

		for j := i + 1; j < len(input); j++ {

			if input[j] < input[minimum] {
				minimum = j

			}
		}

		if minimum != i {
			input[i], input[minimum] = input[minimum], input[i]
		}

	}

	return input

}

func InsertionSort(input []int) []int {

	for i := 1; i < len(input); i++ {

		if input[i] < input[i-1] {

			j := i - 1
			temp := input[i]

			for j >= 0 && input[j] > temp {
				input[j+1] = input[j]
				j--
			}
			input[j+1] = temp
		}

	}

	return input

}

func MergeSort(input []int) []int {

	if len(input) < 2 {
		return input
	}

	leftPart := MergeSort(input[:len(input)/2])
	rightPart := MergeSort(input[len(input)/2:])

	return merge(leftPart, rightPart)

}

func merge(leftPart []int, rightPart []int) []int {

	result := []int{}

	i := 0
	j := 0

	for i < len(leftPart) && j < len(rightPart) {

		if leftPart[i] < rightPart[j] {
			result = append(result, leftPart[i])
			i++
		} else {
			result = append(result, rightPart[j])
			j++
		}
	}

	for ; i < len(leftPart); i++ {
		result = append(result, leftPart[i])
	}

	for ; j < len(rightPart); j++ {
		result = append(result, rightPart[j])
	}

	return result

}

func QuickSort(input []int, left int, right int) {

	if left < right {

		pivot := partion(input, left, right)

		QuickSort(input, left, pivot-1)
		QuickSort(input, pivot+1, right)

	}

}

func partion(input []int, left int, right int) int {

	pivot := input[right]
	i := left - 1

	for j := left; j < right; j++ {
		if input[j] < pivot {
			i++
			input[i], input[j] = input[j], input[i]
		}
	}

	input[i+1], input[right] = input[right], input[i+1]

	return i + 1

}

func QuickSortTnatFails(input []int) []int {

	if len(input) < 2 {
		return input
	}

	index := partionThatFails(input)

	left := QuickSortTnatFails(input[:index])
	right := QuickSortTnatFails(input[index+1:])

	left = append(left, input[0])

	return append(left, right...)

}

func partionThatFails(input []int) int {

	pivot := input[0]

	i := 1
	j := len(input) - 1

	for i < j {

		for (i < len(input)) && (input[i] < pivot) {
			i++
		}

		for (j > 0) && (input[j] > pivot) {
			j--
		}

		if j > i {
			input[i], input[j] = input[j], input[i]
		}
	}

	input[0], input[j] = input[j], input[0]

	return j

}

func theory() {

	//  Pick the most fit sorting algorithms for the given scenarios:

	// #1 - Sort 10 schools around your house by distance:

	/*  The array is almost sorted, as schools are around the house, so
	Insertion sort could be fine for it*/

	// #2 - eBay sorts listings by the current Bid amount:
	/* We have a huge amount of data so the best idea is to use something really
	efficient, dealing with ints: Radix or Counting sort could be ok
	Merge sort is also good*/

	// #3 - Sport scores on ESPN:
	/* A lot of non-numeric data, so QuickSort,
	MergeSort (could be too much of space complexity) or HeapSort could help*/

	// #4 - Massive database (can't fit all into memory) needs to sort through past year's user data:
	/*Looks like a job for a MergeSort*/

	// #5 - Almost sorted Udemy review data needs to update and add 2 new reviews:

	/* Almost sorted data could be sorted either with Insertion sorting or Heap sorting*/

	// #6 - Temperature Records for the past 50 years in Canada:
	/* Numeric data could be effectively sorted with any of these: Radix, Counting - having no decimals ,
	otherwise it would be  MergeSort or QuickSort*/

	// #7 - Large user name database needs to be sorted. Data is very random:

	/* MergeSort or Quicksort are here to help*/

	// #8 - You want to teach sorting for the first time:
	/* Bubble sort and Selection sort could come in handy*/

}
