package luhn

import (
	"testing"
)

func TestIsValidNumber(t *testing.T) {

	digits := []int{4, 9, 9, 2, 7, 3, 9, 8, 7, 1, 6}
	if !IsValidNumber(digits) {
		t.Error("should be valid")
	}
}
