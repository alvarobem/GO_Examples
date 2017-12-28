package main

import (
	"fmt"
)

//Read from keyboard one string and check if it's palindrome
func main() {
	fmt.Println("Write the line:")
	var input string
	fmt.Scanf("%s", &input)
	isPalindrome := checkPalindrome(input)
	fmt.Println(isPalindrome)
}

func checkPalindrome(line string) (isPalindrome bool) {
	lineLen := len(line)
	lineMed := int(lineLen / 2)
	lineLen--
	isPalindrome = true
	i := 0
	for i <= lineMed && isPalindrome {
		if line[i] == line[lineLen-i] {
			i++
		} else {
			isPalindrome = false
		}
	}
	return
}
