package main

import (
	"fmt"
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
	fmt.Println(isPalindrome)
	if !isPalindrome {
		t.Error("checkPalindrome was incorrect expected: true got: false")
	}
}

func TestIsPalindomeEmpry(t *testing.T) {
	isPalindrome := checkPalindrome("")
	if !isPalindrome {
		t.Error("checkPalindrome was incorrect expected: true got: false")
	}
}
