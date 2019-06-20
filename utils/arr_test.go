package utils_test

import (
	"testing"

	"github.com/yoloyi/algorithm/utils"
)

// Test Contain
func TestContain(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	if result, err := utils.Contain(1, arr); err != nil && result {
		t.Error(err)
	}
}

func TestNotContain(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	if _, err := utils.Contain(8, arr); err == nil {
		t.Error(err)
	}
}
