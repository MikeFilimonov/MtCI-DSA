package recursion

func FindFactorialRecursive(number int) int {

	if number <= 1 {
		return 1
	} else {
		return number * FindFactorialRecursive(number-1)
	}

}

func FindFactorialIterative(number int) int {

	if number <= 1 {
		return 0
	}

	result := 1

	for number > 1 {

		result *= number
		number--

	}

	return result

}

func FibonacciIterative(n int) int {

	if n < 0 {
		return 0
	}

	values := []int{}
	values = append(values, 0)
	values = append(values, 1)

	if n < 2 {
		return values[n]
	}

	counter := 2
	for n >= counter {

		values = append(values, values[counter-1]+values[counter-2])
		counter++

	}

	return values[n]

}

func FibonacciRecursive(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}

}

func RevertStringRecursively(input string) string {

	if input == "" {
		return ""
	} else {
		return RevertStringIteratively(input[1:]) + string(input[0])
	}

}

func RevertStringIteratively(input string) string {

	runed := []rune(input)
	for i := 0; i < len(input)/2; i++ {
		runed[i], runed[len(input)-1-i] = runed[len(input)-1-i], runed[i]
	}

	return string(runed)
}
