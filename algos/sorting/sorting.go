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
