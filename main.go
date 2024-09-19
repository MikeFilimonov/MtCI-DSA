package main

import (
	"MtCI-DSA/arrays"
	"MtCI-DSA/hashmaps"
	"MtCI-DSA/linked_lists"
	"fmt"
)

func main() {

	// arraySession()
	// hashMapSession()
	linkedListSession()
	// doubleLinkedListSession()
}

func doubleLinkedListSession() {

	doubleLinkedList := linked_lists.NewDoublyLinkedList(9)
	doubleLinkedList.Append(99)
	doubleLinkedList.Prepend(8)
	doubleLinkedList.Insert(999, 2)
	doubleLinkedList.Delete(0)
	doubleLinkedList.Delete(1)
	doubleLinkedList.ShowList()

}

func linkedListSession() {

	list := linked_lists.New(10)
	list.Append(5)
	list.Append(16)
	list.Prepend(1)
	list.Insert(90, 1)
	list.Insert(11, 10)
	list.ShowList()
	list.Delete(2)
	list.ShowList()
	result := list.Reverse()
	result.ShowList()
	list.ShowList()
	list.BetterReverse()
	list.ShowList()

}

func firstRecurringCharacter(input []int) int {

	buffer := make(map[int]int)

	if len(input) < 2 {
		return -1
	}

	for k, v := range input {

		_, found := buffer[v]
		if found {
			return v
		} else {
			buffer[v] = k
		}

	}

	return -1

}

func hashMapSession() {

	fmt.Println(hashmaps.NaiveHash("ershov"))
	fmt.Println(hashmaps.NaiveHash("choovarsh"))
	h := hashmaps.New(50)
	h.Set("wallet", 10000)
	h.Set("clamp", 8000)
	fmt.Println(h.Get("clamp"))

	fmt.Println(h.Keys())

	// fmt.Println(hashmaps.NaiveHash("sdfsd"))

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
