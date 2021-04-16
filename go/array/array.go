package array

import (
	"errors"
	"fmt"
)

// 扩充传统数组的插入、删除、按照下标随机访问操作；
// 数组中的数据是int类型的
type Array struct {
	length uint
	data   []int
}

func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		length: 0,
		data:   make([]int, capacity, capacity),
	}
}

func (a *Array) Length() uint {
	return a.length
}

func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

func (a *Array) Insert(index uint, value int) error {
	if a.length == uint(cap(a.data)) {
		return errors.New("array is full")
	}

	if a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return nil
}

func (a *Array) InsertToTail(value int) error {
	return a.Insert(a.length, value)
}

func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	value := a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return value, nil
}

func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.length; i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
