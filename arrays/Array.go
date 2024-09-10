package arrays

import (
	"errors"
	"reflect"
)

type Array struct {
	length   int
	capacity int
	data     []interface{}
}

func New(capacity int) Array {
	return Array{
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
