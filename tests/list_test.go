package tests

import (
	"testing"

	structures "github.com/yoloyi/algorithm/data-structures"
)

func TestNewList(t *testing.T) {
	_, error := structures.NewLinearList(5)
	if error != nil {
		t.Errorf(error.Error())
	}
}

func TestAppend(t *testing.T) {
	s, error := structures.NewLinearList(5)
	if error != nil {
		t.Errorf(error.Error())
	}
	if s.Append(1) != nil {
		t.Errorf(error.Error())
	}
	if s.Append(2) != nil {
		t.Errorf(error.Error())
	}

	e, error := s.GetElem(0)
	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 1 {
		t.Errorf("GetElem exists error")
	}

	e, error = s.GetElem(1)
	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 2 {
		t.Errorf("GetElem exists error")
	}
}

func TestClearList(t *testing.T) {
	s, error := structures.NewLinearList(5)
	if error != nil {
		t.Errorf(error.Error())
	}
	if s.Append(1) != nil {
		t.Errorf(error.Error())
	}
	if s.Append(2) != nil {
		t.Errorf(error.Error())
	}
	if s.IsEmpty() {
		t.Errorf("IsEmpty  or Append method exists error")
	}
	s.ClearList()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty  or Append method exists error")
	}
}

func TestListInsert(t *testing.T) {
	s, error := structures.NewLinearList(5)
	if error != nil {
		t.Errorf(error.Error())
	}
	s.ListInsert(0, 1)
	s.ListInsert(0, 2)
	s.ListInsert(1, 3)

	if s.ListLength() != 3 {
		t.Errorf("ListLength exists error")
	}
	e, error := s.GetElem(0)
	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 2 {
		t.Errorf("GetElem exists error")
	}

	e, error = s.GetElem(1)
	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 3 {
		t.Errorf("GetElem exists error")
	}

	e, error = s.GetElem(2)
	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 1 {
		t.Errorf("GetElem exists error")
	}

	e, error = s.ListDelete(1)

	if error != nil {
		t.Errorf(error.Error())
	}
	if e != 3 {
		t.Errorf("ListDelete exists error")
	}
}

func TestIsEmpty(t *testing.T) {
	s, error := structures.NewLinearList(5)
	if error != nil {
		t.Errorf(error.Error())
	}
	if !s.IsEmpty() {
		t.Errorf("IsEmpty exists error")
	}
}
