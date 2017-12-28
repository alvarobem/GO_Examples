package main

import (
	"testing"
)

func TestIsPalindomeCar(t *testing.T) {
	isPalindrome := checkPalindrome("CAR")
	if isPalindrome {
		t.Error("checkPalindrome was incorrect expected: false got: true")
	}
}

func TestIsPalindomeRotator(t *testing.T) {
	isPalindrome := checkPalindrome("Rotator")
	if isPalindrome {
		t.Error("checkPalindrome was incorrect expected: true got: false")
	}
}
