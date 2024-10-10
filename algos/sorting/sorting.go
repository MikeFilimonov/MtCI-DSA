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
