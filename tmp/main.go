package main

import "fmt"

func main() {

	ex2()

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
	for k, v := range collection {

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
