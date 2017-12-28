package main

import (
	"fmt"
	"strings"
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
	line = strings.ToLower(line)
	lineLen := len(line)
	isPalindrome = true
	if lineLen > 0 {
		lineMed := int(lineLen / 2)
		lineLen--
		i := 0
		for i <= lineMed && isPalindrome {
			if line[i] == line[lineLen-i] {
				i++
			} else {
				isPalindrome = false
			}
		}
	}
	return
}
