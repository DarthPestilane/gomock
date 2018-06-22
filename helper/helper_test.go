package helper

import (
	"testing"
)

func TestRandomSlice(t *testing.T) {
	arr := []int{1, 2, 3, 5, 4}
	var val int
	RandomSlice(arr, &val) // without panic
}
