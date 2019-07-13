package structures

import (
	"errors"
	"fmt"
)

type Object interface{}

// LinearList  顺序存储结构线性表
type LinearList struct {
	data    []Object
	length  uint
	maxsize uint
}

// NewLinearList 实例化LinearList
func NewLinearList(maxsize uint) (*LinearList, error) {
	if maxsize == 0 || maxsize < 0 {
		return nil, errors.New("maxsize should uint")
	}
	return &LinearList{maxsize: maxsize, data: make([]Object, maxsize)}, nil
}

// IsEmpty LinearList is empty
func (list LinearList) IsEmpty() bool {
	return list.length == 0
}

// ClearList clear list
func (list *LinearList) ClearList() error {
	if list.length == 0 {
		return errors.New("List is empty")
	}
	list.data = make([]Object, list.maxsize)
	list.length = 0
	return nil
}

// Append append
func (list *LinearList) Append(value Object) error {
	if list.length >= list.maxsize {
		return errors.New("List out of range")
	}
	list.data[list.length] = value
	list.length++
	return nil
}

// GetElem 将线性表第 index 位置，返回
func (list LinearList) GetElem(index uint) (Object, error) {
	if index >= list.maxsize || index < 0 {
		return nil, fmt.Errorf("Index %d out of range", index)
	}
	e := list.data[index]
	return e, nil
}

// ListDelete 删除第 index 位置，返回 index 位置的值
func (list *LinearList) ListDelete(index uint) (Object, error) {
	if index >= list.maxsize || index < 0 {
		return nil, fmt.Errorf("out of range")
	}
	e := list.data[index]
	if e == nil {
		return nil, nil
	}
	list.length--
	for i := index; i < list.maxsize-1; i++ {
		list.data[i], list.data[i+1] = list.data[i+1], list.data[i]
	}
	if list.data[list.maxsize-1] != nil {
		list.data[list.maxsize-1] = nil
	}

	return e, nil
}

// ListInsert 在线性表第 index 个位置，插入新元素e
func (list *LinearList) ListInsert(index uint, e Object) error {
	if index >= list.maxsize || index < 0 || list.length == list.maxsize {
		return fmt.Errorf("out of range")
	}
	list.length++
	if list.data[index] == nil {
		list.data[index] = e
		return nil
	}

	for i := list.maxsize - 1; i > index; i-- {
		list.data[i], list.data[i-1] = list.data[i-1], list.data[i]
	}
	list.data[index] = e
	return nil
}

// ListLength 长度
func (list *LinearList) ListLength() uint {

	return list.length
}
